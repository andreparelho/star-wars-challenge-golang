package service

import (
	"encoding/json"
	"net/http"

	"github.com/andreparelho/star-wars-challenge-golang/planet/model/request"
	"github.com/andreparelho/star-wars-challenge-golang/planet/model/response"
	"github.com/gin-gonic/gin"
)

const apiUrl = "https://swapi.dev/api/planets/"

func GetPlanets(GetPlanetRequest request.GetPlanetRequest, context *gin.Context) response.HandlerGetPlanetsResponse {
	requestApi, error := http.Get(apiUrl)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error to call external API",
		})
		return response.HandlerGetPlanetsResponse{}
	}

	defer requestApi.Body.Close()

	var getPlanetsResponse = response.Data{}
	if error := json.NewDecoder(requestApi.Body).Decode(&getPlanetsResponse); error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error to processing json",
		})
		return response.HandlerGetPlanetsResponse{}
	}

	var planetResponse = checkExistsPlanet(getPlanetsResponse, GetPlanetRequest)
	if (response.HandlerGetPlanetsResponse{}) == planetResponse {
		return response.HandlerGetPlanetsResponse{}
	}

	return planetResponse
}

func checkExistsPlanet(data response.Data, planetRequest request.GetPlanetRequest) response.HandlerGetPlanetsResponse {
	for _, planet := range data.Result {
		if planetRequest.Name == planet.Name {
			return response.HandlerGetPlanetsResponse{
				PlanetName: planet.Name,
				Climate:    planet.Climate,
				Terrain:    planet.Terrain,
				Films:      len(planet.Films),
			}
		}
	}

	return response.HandlerGetPlanetsResponse{}
}
