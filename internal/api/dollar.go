package api

import (
	"net/http"

	"github.com/alejolencinas/dollar-price/internal/scraper"
	"github.com/gin-gonic/gin"
)

func GetDollarPrice(c *gin.Context) {
	price, err := scraper.GetDollarPrice()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"buy":        price.Buy,
		"sell":       price.Sell,
		"fetched_at": price.FetchedAt,
	})
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
