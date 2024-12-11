package handler

import (
	"net/http"

	"github.com/andreparelho/star-wars-challenge-golang/planet/common"
	"github.com/andreparelho/star-wars-challenge-golang/planet/model/request"
	"github.com/andreparelho/star-wars-challenge-golang/planet/service"
	"github.com/gin-gonic/gin"
)

func PostPlanetHandler(context *gin.Context) {
	var postPlanetRequest = request.PostPlanetRequest{}

	var error error = context.BindJSON(&postPlanetRequest)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			common.ERROR_KEY: common.BIND_ERROR_MESSAGE,
		})
		return
	}

	response := service.PostPlanets(postPlanetRequest, context)

	context.JSON(http.StatusOK, response)
}
