package channels

import (
	"fmt"

	pusher "github.com/pusher/pusher-http-go"
)

func Channel_Start(suffix []string, interest string) {
	pusherClient := pusher.Client{
		AppID:   "1355301",
		Key:     "bc6315e1c8ce114b7cb1",
		Secret:  "a220a4346ea0b387b2c5",
		Cluster: "mt1",
		Secure:  true,
	}

	data := map[string]interface{}{
		"key":   interest,
		"value": suffix,
	}
	err := pusherClient.TriggerMulti([]string{"workspace"}, "code", data)
	if err != nil {
		fmt.Println("Failed to trigger", err.Error())
		return
	}
}
