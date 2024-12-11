package main

import (
	"github.com/andreparelho/star-wars-challenge-golang/planet/handler"
	"github.com/gin-gonic/gin"
)

const SERVER_PORT string = ":8080"
const PLANETS_ENDPOINT string = "/planets"

func main() {
	router := gin.Default()

	router.GET(PLANETS_ENDPOINT, handler.GetPlanetHandler)
	router.POST(PLANETS_ENDPOINT, handler.PostPlanetHandler)

	router.Run(SERVER_PORT)
}
