package handler

import (
	"net/http"

	"github.com/andreparelho/star-wars-challenge-golang/planet/model/request"
	"github.com/andreparelho/star-wars-challenge-golang/planet/service"
	"github.com/gin-gonic/gin"
)

func GetPlanetHandler(context *gin.Context) {
	var GetPlanetRequest = request.GetPlanetRequest{}

	if error := context.BindJSON(&GetPlanetRequest); error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "bind json error",
		})
		return
	}

	response := service.GetPlanets(GetPlanetRequest, context)

	context.JSON(http.StatusOK, response)
}
