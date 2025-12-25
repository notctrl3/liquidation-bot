package scanner

import (
	"liquidation-bot/contracts"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
)

func ScanBorrowers(engine *contracts.DSCEngine) []common.Address {
	set := make(map[common.Address]struct{})

	iter, err := engine.FilterCollateralDeposited(
		&bind.FilterOpts{Start: 0, End: nil},
		nil, nil, nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	for iter.Next() {
		set[iter.Event.User] = struct{}{}
	}

	users := make([]common.Address, 0, len(set))
	for user := range set {
		users = append(users, user)
	}
	return users
}
