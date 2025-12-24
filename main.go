// main.go
package main

import (
	"context"
    "fmt"
	"log"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	block, err := client.BlockNumber(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("connected to anvil, block:", block)
}
