// main.go
package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"liquidation-bot/chain"
	"liquidation-bot/engine"
	"liquidation-bot/scanner"
)

func main() {
	client := chain.MustDial("http://127.0.0.1:8545")

	dscEngineAddr := common.HexToAddress("0x6cffa31dd31cF649fb24AC7cEfE87579324bD89c")

	dscEngine := engine.NewDSCEngine(dscEngineAddr, client)

	borrowers := scanner.ScanBorrowers(dscEngine.Raw())

	for _, user := range borrowers {
		hf := dscEngine.HealthFactor(user)
		fmt.Println(user.Hex(), hf)
	}

}
