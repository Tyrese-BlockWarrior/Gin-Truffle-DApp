package services

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/DEMYSTIF/gin-truffle-dapp/lib"
	certTypes "github.com/DEMYSTIF/gin-truffle-dapp/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func InfoService(client *ethclient.Client) (*big.Int, *big.Int, uint64) {
	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	latestBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return networkID, chainID, latestBlock
}

func IssueService(client *ethclient.Client, instance *lib.Cert, newCertificate certTypes.FormCertificate) (*types.Transaction, error) {
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasLimit := uint64(300000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	_, chainID, _ := InfoService(client)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice
	intID, err := strconv.ParseInt(newCertificate.ID, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	trx, err := instance.Issue(auth, big.NewInt(intID), newCertificate.Name, newCertificate.Course, newCertificate.Grade, newCertificate.Date)

	return trx, err
}

func FetchService(instance *lib.Cert, id int64) (certTypes.ReturnCertificate, error) {
	opts := bind.CallOpts{}
	certID := big.NewInt(id)

	result, err := instance.Certificates(&opts, certID)
	return result, err
}
