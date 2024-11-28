package handler

import (
	"log"
	"net/http"

	"github.com/andreparelho/star-wars-challenge-golang/planet/model/request"
	"github.com/andreparelho/star-wars-challenge-golang/planet/service"
	"github.com/gin-gonic/gin"
)

func GetPlanetHandler(context *gin.Context) {
	var createPlanetRequest = request.CreatePlanetRequest{}

	if error := context.BindJSON(&createPlanetRequest); error != nil {
		log.Fatal("bind json error")
		return
	}

	response := service.GetPlanets(createPlanetRequest, context)

	context.JSON(http.StatusOK, response)
}
