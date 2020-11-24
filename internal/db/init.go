package db

import (
	redis "github.com/go-redis/redis/v8"
	minio "github.com/minio/minio-go/v7"
	"gorm.io/gorm"

	"github.com/9d77v/pdc/internal/utils"
)

//环境变量
var (
	DEBUG = utils.GetEnvBool("DEBUG", true)

	dbHost      = utils.GetEnvStr("DB_HOST", "domain.local")
	dbPort      = utils.GetEnvInt("DB_PORT", 5432)
	dbUser      = utils.GetEnvStr("DB_USER", "postgres")
	dbPassword  = utils.GetEnvStr("DB_PASSWORD", "123456")
	dbName      = utils.GetEnvStr("DB_NAME", "pdc")
	TablePrefix = utils.GetEnvStr("DB_PREFIX", "pdc")

	minioAddress         = utils.GetEnvStr("MINIO_ADDRESS", "oss.domain.local:9000")
	secureMinioAddress   = utils.GetEnvStr("SECURE_MINIO_ADDRESS", "oss.domain.local")
	minioAccessKeyID     = utils.GetEnvStr("MINIO_ACCESS_KEY", "minio")
	minioSecretAccessKey = utils.GetEnvStr("MINIO_SECRET_KEY", "minio123")
	OssPrefix            = ""
	SecureOssPrerix      = ""

	redisAddress  = utils.GetEnvStr("REDIS_ADDRESS", "domain.local:6379")
	redisPassword = utils.GetEnvStr("REDIS_PASSWORD", "")
)

var (
	//Gorm global  orm
	Gorm *gorm.DB
	//MinioClient S3 OSS by http
	MinioClient *minio.Client
	//SecureMinioClient S3 OSS by https
	SecureMinioClient *minio.Client
	//RedisClient ..
	RedisClient *redis.Client
)

//角色
const (
	RoleAdmin   = 1
	RoleManager = 2
	RoleUser    = 3
	RoleGuest   = 4
)

//redis前缀
const (
	PrefixUser = "USER"
)

func init() {
	initDB()
	initMinio()
	initRedis()
}

//GetOSSPrefix ..
func GetOSSPrefix(sheme string) string {
	if sheme == "https" {
		return SecureOssPrerix
	}
	return OssPrefix
}
