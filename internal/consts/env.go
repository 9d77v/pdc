package consts

import "github.com/9d77v/pdc/internal/utils"

//环境变量
var (
	DEBUG = utils.GetEnvBool("DEBUG", true)

	JWTtAccessSecret = utils.GetEnvStr("JWT_ACCESS_SECRET", "JWT_ACCESS_SECRET")
	JWTRefreshSecret = utils.GetEnvStr("JWT_REFRESH_SECRET", "JWT_REFRESH_SECRET")
	JWTIssuer        = utils.GetEnvStr("JWT_ISSUER", "domain.local")

	hashSecretUID      = utils.GetEnvStr("PDC_HASH_SECRET_UID", "asdfgh")
	hashUIDLength      = utils.GetEnvInt("PDC_HASH_UID_LENGTH", 10)
	hashSecretDeviceID = utils.GetEnvStr("PDC_HASH_SECRET_DEVICE_ID", "zxcvbn")
	accessKeyLen       = utils.GetEnvInt("PDC_DEVICE_ACCESS_KEY_LENGTH", 12)
	secretKeyLen       = utils.GetEnvInt("PDC_DEVICE_SECRET_KEY_LENGTH", 32)
)

//GetEncodeUID ..
func GetEncodeUID(id uint) string {
	return utils.GenerateHashID(id, hashSecretUID, hashUIDLength)
}

//GetDecodeUID ..
func GetDecodeUID(id string) uint {
	return utils.GetRawID(id, hashSecretUID, hashUIDLength)
}

//GetDeviceAccessKey ..
func GetDeviceAccessKey(id uint) string {
	return utils.GenerateHashID(id, hashSecretDeviceID, accessKeyLen)
}

//GetDeviceSecretKey ..
func GetDeviceSecretKey() string {
	return utils.GenerateSecretKey(secretKeyLen)
}
