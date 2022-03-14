package channels

import (
	pusher "github.com/pusher/pusher-http-go"
)

func Channel_Start(suffix []string, interest string) {

	// channels params
	pusherClient := pusher.Client{
		AppID:   "1355301",
		Key:     "bc6315e1c8ce114b7cb1",
		Secret:  "a220a4346ea0b387b2c5",
		Cluster: "mt1",
		Secure:  true,
	}

	// data object that convey over server to client
	data := map[string]interface{}{
		"key":   interest,
		"value": suffix,
	}

	// trigger event with data object
	err := pusherClient.TriggerMulti([]string{"workspace"}, "code", data)
	if err != nil {
		return
	}
}
