package main

import (
	"log"
	"strings"

	"git.9d77v.me/9d77v/pdc/models"
	"github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	things := make([]*models.Thing, 0)
	videos := make([]*models.Video, 0)
	episodes := make([]*models.Episode, 0)

	err := models.Gorm.Select("pics").Find(&things).Error
	if err != nil {
		log.Panicln("获取tings失败")
	}
	err = models.Gorm.Select("cover").Find(&videos).Error
	if err != nil {
		log.Panicln("获取videos失败")
	}
	err = models.Gorm.Select("url,subtitles").Find(&episodes).Error
	if err != nil {
		log.Panicln("获取episodes失败")
	}
	for _, v := range things {
		newUrls := make([]string, 0)
		for _, v := range v.Pics {
			newUrls = append(newUrls, ReplaceURL(v))
		}
		err := models.Gorm.Model(&models.Thing{}).Update(map[string]interface{}{
			"pics": newUrls,
		}).Error
		if err != nil {
			log.Panicln("更新thing失败", err)
		}
	}
	for _, v := range videos {
		err := models.Gorm.Model(&models.Video{}).Update(map[string]interface{}{
			"cover": ReplaceURL(v.Cover),
		}).Error
		if err != nil {
			log.Panicln("更新thing失败", err)
		}
	}
	for _, v := range episodes {
		newSubtitles := make(postgres.Hstore, 0)
		for k, vv := range v.Subtitles {
			newURL := ReplaceURL(*vv)
			newSubtitles[k] = &newURL
		}
		err := models.Gorm.Model(&models.Episode{}).Update(map[string]interface{}{
			"url":       ReplaceURL(v.URL),
			"subtitles": newSubtitles,
		}).Error
		if err != nil {
			log.Panicln("更新thing失败", err)
		}
	}
}

//ReplaceURL ..
func ReplaceURL(url string) string {
	if url != "" {
		newURL := ""
		arr := strings.Split(url, "/")
		for i := 3; i < len(arr); i++ {
			newURL += "/" + arr[i]
		}
		return newURL
	}
	return url
}