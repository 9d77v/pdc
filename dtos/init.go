package dtos

import "github.com/9d77v/pdc/models"

func getOSSPrefix(sheme string) string {
	if sheme == "https" {
		return models.SecureOssPrerix
	}
	return models.OssPrefix
}
