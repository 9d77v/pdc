package oss

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/9d77v/pdc/internal/utils"
	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

//minio env
var (
	minioAddress         = utils.GetEnvStr("MINIO_ADDRESS", "oss.domain.local:9000")
	secureMinioAddress   = utils.GetEnvStr("SECURE_MINIO_ADDRESS", "oss.domain.local")
	minioAccessKeyID     = utils.GetEnvStr("MINIO_ACCESS_KEY", "minio")
	minioSecretAccessKey = utils.GetEnvStr("MINIO_SECRET_KEY", "minio123")
	OssPrefix            = fmt.Sprintf("http://%s", minioAddress)
	SecureOssPrerix      = fmt.Sprintf("https://%s", secureMinioAddress)
)

var (
	client       *minio.Client
	once         sync.Once
	secureClient *minio.Client
	secureOnce   sync.Once
)

//GetMinioClient get S3 OSS by http
func GetMinioClient() *minio.Client {
	once.Do(func() {
		client = initMinioCLient()
	})
	return client
}

func initMinioCLient() *minio.Client {
	conn, err := minio.New(minioAddress, &minio.Options{
		Creds: credentials.NewStaticV4(minioAccessKeyID, minioSecretAccessKey, ""),
	})
	if err != nil {
		log.Fatalln(err)
	}
	return conn
}

//GetSecureMinioClient get S3 OSS by https
func GetSecureMinioClient() *minio.Client {
	secureOnce.Do(func() {
		secureClient = initSecureMinioClient()
	})
	return secureClient
}

func initSecureMinioClient() *minio.Client {
	conn, err := minio.New(secureMinioAddress, &minio.Options{
		Creds:  credentials.NewStaticV4(minioAccessKeyID, minioSecretAccessKey, ""),
		Secure: true,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return conn
}

//InitMinioBuckets ..
func InitMinioBuckets() {
	preCreatedBuckets := []string{"image", "video", "vtt", "pan", "camera"}
	location := "us-east-1"
	for _, bucketName := range preCreatedBuckets {
		err := GetMinioClient().MakeBucket(context.Background(), bucketName,
			minio.MakeBucketOptions{Region: location})
		if err != nil {
			exists, errBucketExists := GetMinioClient().BucketExists(context.Background(), bucketName)
			if errBucketExists == nil && exists {
				log.Printf("We already own %s\n", bucketName)
			} else {
				log.Fatalln(err)
			}
		} else {
			log.Printf("Successfully created %s\n", bucketName)
		}
		if bucketName != "pan" {
			//mc  policy  set  download  minio/mybucket
			policy := `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":["*"]},"Action": 
["s3:GetObject"],"Resource":["arn:aws:s3:::` + bucketName + `/*"]}]}`
			err := GetMinioClient().SetBucketPolicy(context.Background(), bucketName, policy)
			if err != nil {
				log.Printf("Set bucket:%s policy faield:%v\n", bucketName, err)
			}
		}
	}
}

//GetPresignedURL ..
func GetPresignedURL(ctx context.Context, bucketName, objectName, scheme string) (string, error) {
	var minioClient *minio.Client
	if scheme == "https" {
		minioClient = GetSecureMinioClient()
	} else {
		minioClient = GetMinioClient()
	}
	u, err := minioClient.PresignedPutObject(ctx, bucketName, objectName, 6*time.Hour)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

//GetOSSPrefix ..
func GetOSSPrefix(sheme string) string {
	if sheme == "https" {
		return SecureOssPrerix
	}
	return OssPrefix
}
