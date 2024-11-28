package service

import (
	"encoding/json"
	"net/http"

	"github.com/andreparelho/star-wars-challenge-golang/planet/model/request"
	"github.com/andreparelho/star-wars-challenge-golang/planet/model/response"
	"github.com/gin-gonic/gin"
)

const apiUrl = "https://swapi.dev/api/planets/"

func GetPlanets(createPlanetRequest request.CreatePlanetRequest, context *gin.Context) response.Data {
	responseApi, error := http.Get(apiUrl)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error to call external API",
		})
		return response.Data{}
	}

	defer responseApi.Body.Close()

	var getPlanetsResponse = response.Data{}
	if error := json.NewDecoder(responseApi.Body).Decode(&getPlanetsResponse); error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error to processing json",
		})
		return response.Data{}
	}

	return getPlanetsResponse
}
