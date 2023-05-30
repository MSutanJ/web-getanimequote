package main

import (
	"animequote/handler"
	"github.com/gin-gonic/gin"
)

// type Quote struct {
// 	Anime     string `json:"anime"`
// 	Character string `json:"character"`
// 	Quote     string `json:"quote"`
// }

func main() {
	r := gin.Default()
	r.GET("/quote", handler.GetQuoteHandler)
	r.Run(":8080")

	// r := gin.Default()
	// r.GET("/quote", func(c *gin.Context) {
	// 	quote, err := getAnimeQuote()
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{
	// 			"error": err.Error(),
	// 			"data":  nil,
	// 		})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Found",
	// 		"data": quote,
	// 	} )
	// })
	// r.Run(":8080")
}

// func getAnimeQuote() (*Quote, error) {
// 	resp, err := http.Get("https://animechan.vercel.app/api/random/anime?title=naruto")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("Failed to fetch quote")
// 	}

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var quote Quote
// 	err = json.Unmarshal(body, &quote)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &quote, nil
// }
