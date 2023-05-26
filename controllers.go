package main

import (
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func infoController(ctx *gin.Context, client *ethclient.Client) {
	networkID, chainID, latestBlock := infoService(client)

	ctx.IndentedJSON(http.StatusOK, gin.H{"net_version": networkID, "eth_chainId": chainID, "eth_blockNumber": latestBlock})
}

func issueController(ctx *gin.Context, client *ethclient.Client, instance *Cert) {
	var newCertificate InsertCertificate
	if err := ctx.BindJSON(&newCertificate); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"message": "Bad request"})
		return
	}

	trx, err := issueService(client, instance, newCertificate)
	if err != nil {
		log.Fatal(err)
	}

	ctx.IndentedJSON(http.StatusCreated, trx)
}
