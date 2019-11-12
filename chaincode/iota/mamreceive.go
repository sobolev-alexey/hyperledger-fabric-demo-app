package iota

import (
	"fmt"
	"log"

	"github.com/simia-tech/env"

	"github.com/iotaledger/iota.go/api"
	"github.com/iotaledger/iota.go/mam/v1"
)

func Fetch() {
  var (
  	endpointURL = env.String("ENDPOINT_URL", "https://nodes.devnet.iota.org")
  	mode        = env.String("MODE", "public", env.AllowedValues("public", "private", "restricted"))
  	sideKey     = env.String("SIDE_KEY", "")
  	root 		= env.String("ROOT", "YFJPUERTLJFE9GCDYOKVIACLDSFZV99KUDRYOQZZWNRONRJYJZMOTWSTCCKROWIQJBYSKKECRWXCKIHGZ")
  )

  currentRoot := root.Get()

	cm, err := mam.ParseChannelMode(mode.Get())
	if err != nil {
		log.Fatal(err)
	}

	api, err := api.ComposeAPI(api.HTTPClientSettings{
		URI: endpointURL.Get(),
	})
	if err != nil {
		log.Fatal(err)
	}

	receiver := mam.NewReceiver(api)
	if err := receiver.SetMode(cm, sideKey.Get()); err != nil {
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

    // fmt.Println("messages length", len(messages))
		if len(messages) > 0 {
			currentRoot = nextRoot
      goto loop
		}

  	if len(messages) == 0 {
  		// fmt.Println("no messages found")
  	}
}
