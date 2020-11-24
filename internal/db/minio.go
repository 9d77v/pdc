package db

import (
	"context"
	"fmt"
	"log"

	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func initMinio() {
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
