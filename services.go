package main

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func infoService(client *ethclient.Client) (*big.Int, error) {
	networkID, err := client.NetworkID(context.Background())
	return networkID, err
}