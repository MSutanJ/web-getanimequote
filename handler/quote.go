package handler

import (
	"net/http"
	"animequote/repository"
	"github.com/gin-gonic/gin"
)

func GetQuoteHandler(c *gin.Context) {
	quote, err := repository.GetAnimeQuote()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"data":  nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Quote found",
		"data":    quote,
	})
}