package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rhzx3519/quote/quote"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	pool *quote.Pool
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pool, err = quote.NewPool()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// This is used to avoid cors(request different domains) problem from the client
func corsHeader(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")

	// When React calls an API, it first sends an OPTIONS request to detect if the API available
	// So return 204 whenever receive an OPTIONS request to avoid CORS error
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
}

func main() {
	{
		pool.Start()
	}
	r := gin.Default()
	r.Use(corsHeader)

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		v1.GET("/quote", func(c *gin.Context) {
			symbols := c.Query("symbols")
			exchange := c.Query("exchange")

			var items []*quote.Item
			for _, symbol := range strings.Split(symbols, ",") {
				item, err := pool.GetQuote(strings.TrimSpace(symbol), exchange)
				if err == nil {
					items = append(items, item)
				}
			}
			var resp = quote.QuoteResp{
				Data:             quote.Data{Items: items, ItemsSize: len(items)},
				ErrorCode:        0,
				ErrorDescription: "",
			}
			c.JSON(http.StatusOK, resp)
		})
	}

	port := fmt.Sprintf(":%v", os.Getenv("PORT"))
	r.Run(port)
}
