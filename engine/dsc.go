package engine

import (
	"liquidation-bot/contracts"

	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type DSCEngine struct {
	addr common.Address
	raw  *contracts.DSCEngine
}

func NewDSCEngine(addr common.Address, client *ethclient.Client) *DSCEngine {
	raw, err := contracts.NewDSCEngine(addr, client)
	if err != nil {
		log.Fatal(err)
	}

	return &DSCEngine{
		addr: addr,
		raw:  raw,
	}
}

func (e *DSCEngine) HealthFactor(user common.Address) *big.Int {
	hf, err := e.raw.GetHealthFactor(nil, user)
	if err != nil {
		log.Panicln("hf err:", err)
		return nil
	}
	return hf
}

func (e *DSCEngine) Raw() *contracts.DSCEngine {
	return e.raw
}
