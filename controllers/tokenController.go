package controllers

import (
	"encoding/json"
	"net/http"
	"testhex10/rest_api/models"

	"github.com/gin-gonic/gin"
)

func GetTokenByWallet(c *gin.Context) {
	var tokens []models.Token
	wallet_address := c.Param("wallet_address")
	models.DB.Where("owner = ?", wallet_address).Find(&tokens)
	if len(tokens) < 1 {
		c.JSON(http.StatusNotFound, gin.H{"message": "data not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": tokens})
}

func PostTokenByWallet(c *gin.Context) {
	wallet_address, wallet_address_status := c.GetQuery("wallet_address")
	if wallet_address_status == false {
		c.JSON(http.StatusBadRequest, gin.H{"message": "wallet_address is required"})
		return
	}
	baseURL := "https://api-mainnet.magiceden.dev/v2/wallets/" + wallet_address + "/tokens?offset=0&limit=10&listStatus=listed"
	request, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "request failed"})
		return
	}
	var client = &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer response.Body.Close()
	var items []models.Token
	err = json.NewDecoder(response.Body).Decode(&items)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "decode failed"})
		return
	}

	for _, data := range items {
		models.DB.Create(data)
	}

	c.JSON(204, gin.H{"message": "data success been created"})
}

type inputDelete struct {
	TokenMintAddress string `json:"token_mint_address" binding:"required"`
}

func DeleteToken(c *gin.Context) {
	var input inputDelete
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var token models.Token
	err = models.DB.Where("mint_address = ?", input.TokenMintAddress).First(&token).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data not faund"})
		return
	}
	models.DB.Where("mint_address = ?", input.TokenMintAddress).Delete(&token)
	c.JSON(http.StatusOK, gin.H{"message": "delete success token"})
}
