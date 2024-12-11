package handler

import (
	"net/http"

	"github.com/andreparelho/star-wars-challenge-golang/planet/common"
	"github.com/andreparelho/star-wars-challenge-golang/planet/service"
	"github.com/gin-gonic/gin"
)

func GetPlanetHandler(context *gin.Context) {
	planetsService := &service.GetAllPlanetsService{
		Client: &service.RealHTTPClient{},
	}

	response, error := planetsService.GetAllPlanets(context)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			common.ERROR_KEY: common.HANDLER_ERROR_MESSAGE,
		})
	}

	context.JSON(http.StatusOK, response)
}
