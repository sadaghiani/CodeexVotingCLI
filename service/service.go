package service

import (
	"codeex/app/vote"
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Eth is the ehereum service
type Eth struct {
	client          *ethclient.Client
	chainID         *big.Int
	contractAddress common.Address
	instance        *vote.Vote
	keystore        *keystore.KeyStore
}

// NewEth create new eth service
func NewEth(providerAddress string, contractAddress string) (*Eth, error) {
	c, err := ethclient.Dial(providerAddress)
	if err != nil {
		return nil, err
	}

	chainID, err := c.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	ins, err := vote.NewVote(common.HexToAddress(contractAddress), c)

	return &Eth{
		client:          c,
		chainID:         chainID,
		contractAddress: common.HexToAddress(contractAddress),
		instance:        ins,
		keystore:        keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP),
	}, nil

}
