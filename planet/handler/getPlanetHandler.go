package handler

import (
	"net/http"

	"github.com/andreparelho/star-wars-challenge-golang/planet/model/request"
	"github.com/andreparelho/star-wars-challenge-golang/planet/service"
	"github.com/gin-gonic/gin"
)

func GetPlanetHandler(context *gin.Context) {
	var getPlanetRequest = request.GetPlanetRequest{}

	if error := context.BindJSON(&getPlanetRequest); error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "bind json error",
		})
		return
	}

	response := service.GetPlanets(getPlanetRequest, context)

	context.JSON(http.StatusOK, response)
}
