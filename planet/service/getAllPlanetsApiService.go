package service

import (
	"encoding/json"
	"net/http"

	"github.com/andreparelho/star-wars-challenge-golang/planet/common"
	"github.com/andreparelho/star-wars-challenge-golang/planet/model/response"
	"github.com/gin-gonic/gin"
)

const QUERY_PARAM string = "?page="

func GetAllPlanets(context *gin.Context) response.HandlerGetPlanetsResponse {
	var page string = context.Query("page")

	requestApi, error := http.Get(common.API_URL + QUERY_PARAM + page)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			common.ERROR_KEY: common.CALL_API_ERROR_MESSAGE,
		})
		return response.HandlerGetPlanetsResponse{}
	}

	defer requestApi.Body.Close()

	var getPlanetsResponse = response.Data{}
	error = json.NewDecoder(requestApi.Body).Decode(&getPlanetsResponse)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			common.ERROR_KEY: common.JSON_PROCESS_ERROR_MESSAGE,
		})
		return response.HandlerGetPlanetsResponse{}
	}

	var responseHandler = response.HandlerGetPlanetsResponse{
		Planets: getPlanetsResponse.Result,
	}

	return responseHandler
}
