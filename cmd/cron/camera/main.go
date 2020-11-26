package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/robfig/cron/v3"

	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/device-service/models"
)

func main() {
	var isDryRun = flag.Bool("dryRun", false, "Input Your Command")
	var idStr = flag.String("ids", "", "Input camera ids")
	var yesterdayStr = flag.String("yesterdayStr", "", "Input camera yesterdayStr")
	flag.Parse()

	if *isDryRun {
		ids := getIDs(idStr)
		synthesizeJPGsIntoMP4(ids, ptrs.String(yesterdayStr))
	} else {
		cr := cron.New(cron.WithSeconds())
		id, err := cr.AddFunc("0 0 1 * * *", func() {
			synthesizeJPGsIntoMP4([]uint{}, "")
		})
		if err != nil {
			log.Println("cron add func error,entityID:", id, " error:", err)
		}
		cr.Start()
		defer cr.Stop()

		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)
		<-interrupt
	}
}

func getIDs(idStr *string) []uint {
	ids := make([]uint, 0, 0)
	if ptrs.String(idStr) != "" {
		idArr := strings.Split(ptrs.String(idStr), ",")
		for _, v := range idArr {
			id, err := strconv.Atoi(v)
			if id != 0 && err == nil {
				ids = append(ids, uint(id))
			}
		}
	}
	return ids
}

func synthesizeJPGsIntoMP4(ids []uint, yesterdayStr string) {
	if len(ids) == 0 {
		ids = getCameraIDs()
	}
	if len(ids) > 0 {
		if yesterdayStr == "" {
			yesterdayStr = getYesterdayStr()
		}
		for _, id := range ids {
			pictureTmpDir := getCameraPictureTmpDir(id, yesterdayStr)
			err := renamePictureFileNames(pictureTmpDir)
			if err != nil {
				log.Println("renamePictureFileNames failed: ", err)
				continue
			}
			videoFilePath := getVideoPath(id, yesterdayStr)
			err = generateVideo(pictureTmpDir, videoFilePath)
			if err != nil {
				log.Println("generate Video failed: ", err)
				continue
			}
			err = saveVideoPath(id, yesterdayStr, "/camera/"+videoFilePath)
			if err != nil {
				log.Println("saveVideoPath failed: ", err)
			}
		}
	}
}

func getCameraIDs() []uint {
	ids := make([]uint, 0, 0)
	devices := make([]*models.Device, 0, 0)
	err := db.GetDB().Select(db.TablePrefix + "_device.id").
		Where(db.TablePrefix + "_device_model.device_type = 1").
		Joins("JOIN " + db.TablePrefix + "_device_model ON " + db.TablePrefix + "_device_model.id = " +
			db.TablePrefix + "_device.device_model_id").Find(&devices).Error
	if err != nil {
		return ids
	}
	for _, v := range devices {
		ids = append(ids, v.ID)
	}
	return ids
}

func getYesterdayStr() string {
	now := time.Now()
	return now.Add(-24 * time.Hour).Format("2006-01-02")
}

func getCameraPictureTmpDir(id uint, day string) string {
	return fmt.Sprintf("picture/%d/%s/tmp/", id, day)
}

//renamePictureFileNames change file name from timestamp.jpg to %04d.jpg , easy for ffmpeg processing
//example: before: 1606214040.jpg after:  0000.jpg
func renamePictureFileNames(pictureTmpDir string) error {
	rd, err := ioutil.ReadDir(pictureTmpDir)
	if err != nil {
		return err
	}
	for i, fi := range rd {
		fmt.Println(fi.Name())
		os.Rename(pictureTmpDir+fi.Name(), pictureTmpDir+getNewFileName(i)+".jpg")
	}
	return nil
}

func getNewFileName(i int) string {
	prefix := ""
	c := 0
	t := i
	for t > 0 {
		t /= 10
		c++
	}
	if c == 0 {
		c = 1
	}
	for k := 0; k < 4-c; k++ {
		prefix += "0"
	}
	return prefix + strconv.FormatInt(int64(i), 10)
}

func getVideoPath(id uint, day string) string {
	return fmt.Sprintf("picture/%d/%s/%s.mp4", id, day, day)
}

func generateVideo(pictureTmpDir, videoFilePath string) error {
	cmd := exec.Command("ffmpeg",
		"-f", "image2",
		"-threads", "2",
		"-i", pictureTmpDir+"%04d.jpg",
		"-vcodec", "libx264",
		"-r", "10",
		videoFilePath)
	return cmd.Run()
}

func saveVideoPath(deviceID uint, date, videoURL string) error {
	return db.GetDB().Save(models.NewCameraTimeLapseVideo(deviceID, date, videoURL)).Error
}
