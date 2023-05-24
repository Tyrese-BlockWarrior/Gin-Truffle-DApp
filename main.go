package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func home(ctx *gin.Context, client *ethclient.Client) {
	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	
	ctx.IndentedJSON(http.StatusOK, gin.H{"net_version": networkID})
}

func main() {
	file, err := os.Open("details.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	details := Details{}
	err = decoder.Decode(&details)
	if err != nil {
		log.Fatal(err)
	}

	contract := fmt.Sprintf("Contract: %s", details.Contract)
	fmt.Println(contract)
	contractAddress := common.HexToAddress(details.Contract)
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	instance, err := NewCert(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	_ = instance

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		home(ctx, client)
	})
	router.Run("localhost:8080")
}
