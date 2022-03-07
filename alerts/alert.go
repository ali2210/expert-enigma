package alerts

import (
	"fmt"

	pushnotifications "github.com/pusher/push-notifications-go"
)

var counter int64 = 0

const instanceId = "10622845-f1ad-4a5a-b078-7ca7e39b770a"

const secretKey = "0595D60941F36A0B21B1474F630395D8238E3355DCD3A2FC91E37A008148B8FD"

func Watchpoint(message_interface map[string]interface{}, interest string) bool {
	beamsClient, err := pushnotifications.New(instanceId, secretKey)
	if err != nil {
		fmt.Println(" Error creating notify object", err.Error())
		return false
	}

	publishRequest := map[string]interface{}{

		"web": map[string]interface{}{
			"notification": message_interface,
		},
	}

	counter += 1
	_, err = beamsClient.PublishToInterests([]string{interest}, publishRequest)

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true

}

func TAlerts() int64 {
	return counter
}

func SetAlerts(alerts int64) {
	counter = alerts
}
