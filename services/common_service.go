package services

import (
	"time"

	"github.com/9d77v/pdc/models"
)

//CommonService ..
type CommonService struct {
}

//PresignedURL ..
func (s CommonService) PresignedURL(bucketName, objectName string) (string, error) {
	u, err := models.MinioClient.PresignedPutObject(bucketName, objectName, 1*time.Hour)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
