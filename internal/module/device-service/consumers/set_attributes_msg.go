package consumers

import (
	"log"

	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/device-service/models"
	"github.com/9d77v/pdc/pkg/iot/sdk/pb"
)

type setAttributesMsg struct {
	*deviceUpMsg
}

func getSetAttributesMsg() *setAttributesMsg {
	return &setAttributesMsg{
		getDeviceMsg(),
	}
}

func (m *setAttributesMsg) handleMsg(deviceMsg *pb.DeviceUpMsg) {
	attributeMsg := deviceMsg.GetSetAttributesMsg()
	if attributeMsg.AttributeMap == nil {
		return
	}
	for k, v := range attributeMsg.AttributeMap {
		err := db.GetDB().Model(&models.Attribute{}).
			Where("id=?", k).
			Update("value", v).Error
		if err != nil {
			log.Printf("update attribute failed,id:%d,value:%s\n", deviceMsg.DeviceId, v)
		}
	}
}
