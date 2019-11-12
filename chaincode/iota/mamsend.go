package iota

import (
	"fmt"
	"log"
	"encoding/json"
	"strings"
	"io"

	"github.com/simia-tech/env"
	"github.com/iotaledger/iota.go/mam/v1"
)

var (
	// endpointURL = env.String("ENDPOINT_URL", "https://nodes.devnet.iota.org")
	// mamseed     = env.String("SEED", GenerateSeed())
	// mammwm      = env.Int("MWM", 9)
	mode        = env.String("MODE", "public", env.AllowedValues("public", "private", "restricted"))
	sideKey     = env.String("SIDE_KEY", "")
)

func PublishAndStoreState(messages string, useTransmitter bool) {
	var t *mam.Transmitter = nil
	if useTransmitter == true {
		seedFromStorage := Read("seed")
		mamStateFromStorage := Read("mamstate")
		mamState := StringToMamState(mamStateFromStorage)
		t = ReconstructTransmitter(seedFromStorage, mamState)
	}

	transmitter, seed := Publish(messages, t)
	channel := transmitter.Channel()

	Store(MamStateToString(channel), "mamstate")
	Store(seed, "seed")
}

func Publish(messages string, t *mam.Transmitter) (*mam.Transmitter, string) {

	reader := strings.NewReader(messages)
	dec := json.NewDecoder(reader)


	cm, err := mam.ParseChannelMode(mode.Get())
	if err != nil {
		log.Fatal(err)
	}


	api := GetApi()
	if api == nil {
		log.Fatal(err)
	}

	transmitter, seed := GetTransmitter(t, api, cm)

	for {
	    // Read one JSON object and store it in a map.
	    var m interface{}
	    if err := dec.Decode(&m); err == io.EOF {
	        break
	    } else if err != nil {
	        log.Fatal(err)
	    }

		var js []byte
		js, err := json.Marshal(m)
		if err != nil {
			log.Println(err)
		}

		message := string(js)
		fmt.Printf("transmit message %q to %s channel...\n", message, cm)
		root, err := transmitter.Transmit(message) //, "HYPERLEDGER")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("transmitted to root %q\n", root)

	}

	return transmitter, seed
}
