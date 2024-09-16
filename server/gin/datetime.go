package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetDatetime(c *gin.Context) {
	t := time.Now().Format(time.RFC3339)
	c.String(http.StatusOK, t)
}

func StartServer() error {
	router := gin.Default()
	router.GET("/datetime", GetDatetime)
	return router.Run(":8000")
}

func main() {
	if err := StartServer(); err != nil {
		log.Fatal(err)
	}
}
