package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/DEMYSTIF/gin-truffle-dapp/lib"
	"github.com/DEMYSTIF/gin-truffle-dapp/services"
	"github.com/DEMYSTIF/gin-truffle-dapp/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func InfoController(ctx *gin.Context, client *ethclient.Client) {
	networkID, chainID, latestBlock := services.InfoService(client)

	ctx.IndentedJSON(http.StatusOK, gin.H{"net_version": networkID, "eth_chainId": chainID, "eth_blockNumber": latestBlock})
}

func IssueController(ctx *gin.Context, client *ethclient.Client, instance *lib.Cert) {
	var newCertificate types.InsertCertificate
	if err := ctx.BindJSON(&newCertificate); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"message": "Bad request"})
		return
	}

	trx, err := services.IssueService(client, instance, newCertificate)
	if err != nil {
		log.Fatal(err)
	}

	ctx.IndentedJSON(http.StatusCreated, trx)
}

func FetchController(ctx *gin.Context, instance *lib.Cert) {
	param := ctx.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	result, err := services.FetchService(instance, id)
	if err != nil {
		ctx.AbortWithStatus(400)
	}

	ctx.IndentedJSON(http.StatusOK, result)
}
