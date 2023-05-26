package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

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

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		infoController(ctx, client)
	})
	router.POST("/issue", func(ctx *gin.Context) {
		issueController(ctx, client, instance)
	})
	router.GET(("/fetch/:id"), func(ctx *gin.Context) {
		fetchController(ctx, instance)
	})
	router.Run("localhost:8080")
}
