package services

import (
	"context"
	"time"

	"github.com/9d77v/pdc/models"
	minio "github.com/minio/minio-go/v7"
)

//CommonService ..
type CommonService struct {
}

//PresignedURL ..
func (s CommonService) PresignedURL(ctx context.Context, scheme, bucketName, objectName string) (string, error) {
	var minioClient *minio.Client
	if scheme == "https" {
		minioClient = models.SecureMinioClient
	} else {
		minioClient = models.MinioClient
	}
	u, err := minioClient.PresignedPutObject(ctx, bucketName, objectName, 6*time.Hour)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
