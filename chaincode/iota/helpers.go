package iota

import (
  "crypto/rand"
  "math/big"
  "log"
  "fmt"
  "encoding/json"

  "github.com/iotaledger/iota.go/mam/v1"
  "github.com/iotaledger/iota.go/consts"
	"github.com/iotaledger/iota.go/api"
	"github.com/iotaledger/iota.go/pow"
  "github.com/iotaledger/iota.go/trinary"
)

const endpoint = "https://nodes.devnet.iota.org"

// difficulty of the proof of work required to attach a transaction on the tangle
const mwm = 9

// how many milestones back to start the random walk from
const depth = 3

func GenerateSeed() string {
  seed := ""
  alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ9"

  for i := 0; i < 81; i++ {
    n, err := rand.Int(rand.Reader, big.NewInt(27))
    if err != nil {
      log.Fatal(err)
    }
    seed += string(alphabet[n.Int64()])
  }
  return seed
}

func GetTransmitter(t *mam.Transmitter, mode string, sideKey string) (*mam.Transmitter, string) {
  cm, err := mam.ParseChannelMode(mode)
	if err != nil {
		log.Fatal(err)
	}

	api := GetApi()
	if api == nil {
		log.Fatal(err)
  }
  
  switch {
    case t != nil:
      return t, ""
    default:
      seed := GenerateSeed()
      transmitter := mam.NewTransmitter(api, seed, uint64(mwm), consts.SecurityLevelMedium)
      if err := transmitter.SetMode(cm, sideKey); err != nil {
        log.Fatal(err)
      }
      return transmitter, seed
  }
}

func GetApi() *api.API {
  _, powFunc := pow.GetFastestProofOfWorkImpl()

	api, err := api.ComposeAPI(api.HTTPClientSettings{
		URI:                  endpoint,
		LocalProofOfWorkFunc: powFunc,
  })
  if err != nil {
    log.Fatal(err)
    return nil
  }

  return api
}

func ReconstructTransmitter(seed trinary.Trytes, channel *mam.Channel) *mam.Transmitter {
  api := GetApi()

  if api != nil {
    transmitter := mam.NewTransmitterWithChannel(api, seed, uint64(mwm), channel)
    return transmitter
  }

  return nil
}

func MamStateToString(channel *mam.Channel) string {
	jsonChannel, err := json.Marshal(channel)
	if err != nil {
		fmt.Println(err)
	}

	return string(jsonChannel)
}

func StringToMamState(mamstate string) *mam.Channel {
	var channel *mam.Channel
	err := json.Unmarshal([]byte(mamstate), &channel)
	if err != nil {
		fmt.Println("error:", err)
	}
	return channel
}