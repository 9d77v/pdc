package consts

import "log"

func init() {
	if accessKeyLen < 10 {
		log.Panicln("PDC_DEVICE_ACCESS_KEY_LENGTH should longer than or equal to 10")
	}
	if secretKeyLen < 20 {
		log.Panicln("PDC_DEVICE_SECRET_KEY_LENGTH should longer than or equal to 20")
	}
}
