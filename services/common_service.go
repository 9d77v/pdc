package services

import (
	"strings"
	"time"

	"github.com/9d77v/pdc/models"
)

//CommonService ..
type CommonService struct {
}

//PresignedURL ..
func (s CommonService) PresignedURL(scheme, bucketName, objectName string) (string, error) {
	u, err := models.MinioClient.PresignedPutObject(bucketName, objectName, 1*time.Hour)
	if err != nil {
		return "", err
	}
	presignedURL := u.String()
	if scheme == "http" {
		u.Scheme = scheme
		presignedURL = strings.Replace(u.String(), models.MinioAddress, models.InternalMinioAddress, 1)
	}
	return presignedURL, nil
}
