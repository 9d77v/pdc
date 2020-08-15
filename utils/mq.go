package utils

import "log"

//AckHandler for nats ack
func AckHandler(ackedNuid string, err error) {
	if err != nil {
		log.Printf("Warning: error publishing msg id %s: %v\n", ackedNuid, err.Error())
	} else {
		log.Printf("Received ack for msg id %s\n", ackedNuid)
	}
}
