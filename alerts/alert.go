package alerts

import (
	pushnotifications "github.com/pusher/push-notifications-go"
)

// events records
var counter int64 = 0

// notification api credentials
const instanceId = "10622845-f1ad-4a5a-b078-7ca7e39b770a"

const secretKey = "0595D60941F36A0B21B1474F630395D8238E3355DCD3A2FC91E37A008148B8FD"

func Watchpoint(message_interface map[string]interface{}, interest string) bool {

	// valid beam channel credentials
	beamsClient, err := pushnotifications.New(instanceId, secretKey)
	if err != nil {
		return false
	}

	// notification message
	publishRequest := map[string]interface{}{

		"web": map[string]interface{}{
			"notification": message_interface,
		},
	}

	// events triggers
	counter += 1

	// if beam return nil then throw exception
	_, err = beamsClient.PublishToInterests([]string{interest}, publishRequest)

	var valid bool
	if err != nil {
		valid = false
		return valid
	}

	valid = true
	return valid

}

// count number of events triggered
func TAlerts() int64 {
	return counter
}

// reset events trigger
func SetAlerts(alerts int64) {
	counter = alerts
}
