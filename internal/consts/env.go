package consts

import (
	"github.com/9d77v/go-pkg/env"
	"github.com/9d77v/pdc/internal/utils"
)

//环境变量
var (
	DEBUG = env.GetEnvBool("DEBUG", true)

	JWTtAccessSecret = env.GetEnvStr("JWT_ACCESS_SECRET", "JWT_ACCESS_SECRET")
	JWTRefreshSecret = env.GetEnvStr("JWT_REFRESH_SECRET", "JWT_REFRESH_SECRET")
	JWTIssuer        = env.GetEnvStr("JWT_ISSUER", "domain.local")

	hashSecretDeviceID = env.GetEnvStr("PDC_HASH_SECRET_DEVICE_ID", "zxcvbn")
	accessKeyLen       = env.GetEnvInt("PDC_DEVICE_ACCESS_KEY_LENGTH", 12)
	secretKeyLen       = env.GetEnvInt("PDC_DEVICE_SECRET_KEY_LENGTH", 32)
)

//GetDeviceAccessKey ..
func GetDeviceAccessKey(id uint) string {
	return utils.GenerateHashID(id, hashSecretDeviceID, accessKeyLen)
}

//GetDeviceSecretKey ..
func GetDeviceSecretKey() string {
	return utils.GenerateSecretKey(secretKeyLen)
}
