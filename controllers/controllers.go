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

func IndexPageController(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Certificate DApp",
	})
}

func IssuePageController(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "issue.html", gin.H{
		"title": "Certificate DApp", "form": "", "message": "is-hidden", "issuedID": "",
	})
}

func IssueController(ctx *gin.Context, client *ethclient.Client, instance *lib.Cert) {
	var newCertificate types.FormCertificate

	if err := ctx.ShouldBind(&newCertificate); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"message": "Bad request"})
		return
	}

	trx, err := services.IssueService(client, instance, newCertificate)
	if err != nil {
		log.Fatal(err)
	}

	_ = trx
	ctx.HTML(http.StatusCreated, "issue.html", gin.H{
		"title": "Certificate DApp", "form": "is-hidden", "message": "", "issuedID": newCertificate.ID,
	})
}

func FetchController(ctx *gin.Context, instance *lib.Cert) {
	param := ctx.Query("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	result, err := services.FetchService(instance, id)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"message": "Bad request"})
		return
	}

	ctx.HTML(http.StatusOK, "fetch.html", gin.H{
		"title": "Certificate DApp", "id": id, "name": result.Name, "course": result.Course, "grade": result.Grade, "date": result.Date,
	})
}
