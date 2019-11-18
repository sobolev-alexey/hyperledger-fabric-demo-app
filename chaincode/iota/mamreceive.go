package iota

import (
	"fmt"
	"log"

	"github.com/iotaledger/iota.go/api"
	"github.com/iotaledger/iota.go/mam/v1"
)

func Fetch(endpointURL string, root, string, mode string, sideKey string) {
	// var endpointURL = "https://nodes.devnet.iota.org"
	// var mode        = "public"
	// var sideKey     = ""
	// var root 		= "YFJPUERTLJFE9GCDYOKVIACLDSFZV99KUDRYOQZZWNRONRJYJZMOTWSTCCKROWIQJBYSKKECRWXCKIHGZ"

  	currentRoot := root

	cm, err := mam.ParseChannelMode(mode)
	if err != nil {
		log.Fatal(err)
	}

	api, err := api.ComposeAPI(api.HTTPClientSettings{
		URI: endpointURL,
	})
	if err != nil {
		log.Fatal(err)
	}

	receiver := mam.NewReceiver(api)
	if err := receiver.SetMode(cm, sideKey); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("receive root %q from %s channel...\n", currentRoot, cm)

	loop:
		nextRoot, messages, err := receiver.Receive(currentRoot)
		if err != nil {
			log.Fatal(err)
		}

		for _, message := range messages {
			fmt.Println(message)
		}

		if len(messages) > 0 {
			currentRoot = nextRoot
      		goto loop
		}
}
