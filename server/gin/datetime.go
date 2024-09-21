package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetDatetime(c *gin.Context) {
	t := time.Now().Format(time.RFC3339)

	acceptHeader := c.GetHeader("Accept")
	if strings.Contains(acceptHeader, "application/json") {
		c.JSON(http.StatusOK, map[string]string{"datetime": t})
	} else {
		c.String(http.StatusOK, t)
	}
}

func StartServer() error {
	router := gin.Default()
	router.GET("/datetime", GetDatetime)
	return router.Run(":9000")
}

func main() {
	if err := StartServer(); err != nil {
		log.Fatal(err)
	}
}
