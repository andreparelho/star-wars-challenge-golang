package handler

import (
	"net/http"

	"github.com/andreparelho/star-wars-challenge-golang/planet/service"
	"github.com/gin-gonic/gin"
)

func GetPlanetHandler(context *gin.Context) {

	response := service.GetPlanets()

	context.JSON(http.StatusOK, response)
}
