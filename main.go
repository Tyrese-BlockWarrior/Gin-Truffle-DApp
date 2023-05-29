package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DEMYSTIF/gin-truffle-dapp/controllers"
	"github.com/DEMYSTIF/gin-truffle-dapp/lib"
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

	contract := os.Getenv("CONTRACT_ADDRESS")
	printContract := fmt.Sprintf("Contract: %s", contract)
	fmt.Println(printContract)
	contractAddress := common.HexToAddress(contract)
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	instance, err := lib.NewCert(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.Static("/static", "./assets")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(ctx *gin.Context) {
		controllers.IndexPageController(ctx)
	})
	router.GET("/issue", func(ctx *gin.Context) {
		controllers.IssuePageController(ctx)
	})
	router.POST("/issue", func(ctx *gin.Context) {
		controllers.IssueController(ctx, client, instance)
	})
	router.GET(("/fetch"), func(ctx *gin.Context) {
		controllers.FetchController(ctx, instance)
	})
	router.GET("/info", func(ctx *gin.Context) {
		controllers.InfoController(ctx, client)
	})
	router.Run("localhost:8080")
}
