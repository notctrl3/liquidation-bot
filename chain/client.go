package chain

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func MustDial(url string) *ethclient.Client {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
