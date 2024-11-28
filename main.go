package main

import (
	"fmt"

	"github.com/andreparelho/star-wars-challenge-golang/planet/handler"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting server... ")

	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.GET("/ping", handler.Pong)
	router.GET("/planets", handler.GetPlanetHandler)

	router.Run(":8080")
}
