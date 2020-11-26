package oss

import (
	"context"
	"fmt"
	"log"
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
	OssPrefix            = ""
	SecureOssPrerix      = ""
)

var (
	//MinioClient S3 OSS by http
	MinioClient *minio.Client
	//SecureMinioClient S3 OSS by https
	SecureMinioClient *minio.Client
)

func init() {
	accessKeyID := minioAccessKeyID
	secretAccessKey := minioSecretAccessKey

	OssPrefix = fmt.Sprintf("http://%s", minioAddress)
	SecureOssPrerix = fmt.Sprintf("https://%s", secureMinioAddress)
	var err error
	MinioClient, err = minio.New(minioAddress, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		log.Fatalln(err)
	}
	SecureMinioClient, err = minio.New(secureMinioAddress, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: true,
	})
	preCreatedBuckets := []string{"image", "video", "vtt", "pan", "camera"}
	location := "us-east-1"
	for _, bucketName := range preCreatedBuckets {
		err = MinioClient.MakeBucket(context.Background(), bucketName,
			minio.MakeBucketOptions{Region: location})
		if err != nil {
			exists, errBucketExists := MinioClient.BucketExists(context.Background(), bucketName)
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
			err := MinioClient.SetBucketPolicy(context.Background(), bucketName, policy)
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
		minioClient = SecureMinioClient
	} else {
		minioClient = MinioClient
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
