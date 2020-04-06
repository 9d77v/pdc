package services

import (
	"time"

	"git.9d77v.me/9d77v/pdc/models"
)

//CommonService ..
type CommonService struct {
}

//PresignedURL ..
func (s CommonService) PresignedURL(bucketName, objectName string) (string, error) {
	u, err := models.MinioClient.PresignedPutObject(bucketName, objectName, 10*time.Minute)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
