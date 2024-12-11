package service

import (
	"encoding/json"
	"net/http"

	"github.com/andreparelho/star-wars-challenge-golang/planet/common"
	"github.com/andreparelho/star-wars-challenge-golang/planet/model/request"
	"github.com/andreparelho/star-wars-challenge-golang/planet/model/response"
	"github.com/gin-gonic/gin"
)

func PostPlanets(postPlanetRequest request.PostPlanetRequest, context *gin.Context) response.HandlerPostPlanetsResponse {
	requestApi, error := http.Get(common.API_URL)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			common.ERROR_KEY: common.CALL_API_ERROR_MESSAGE,
		})
		return response.HandlerPostPlanetsResponse{}
	}

	defer requestApi.Body.Close()

	var getPlanetsResponse = response.Data{}
	error = json.NewDecoder(requestApi.Body).Decode(&getPlanetsResponse)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			common.ERROR_KEY: common.JSON_PROCESS_ERROR_MESSAGE,
		})
		return response.HandlerPostPlanetsResponse{}
	}

	var planetResponse = checkExistsPlanet(getPlanetsResponse, postPlanetRequest)
	if (response.HandlerPostPlanetsResponse{}) == planetResponse {
		return response.HandlerPostPlanetsResponse{}
	}

	return planetResponse
}

func checkExistsPlanet(data response.Data, planetRequest request.PostPlanetRequest) response.HandlerPostPlanetsResponse {
	for _, planet := range data.Result {
		if planetRequest.Name == planet.Name {
			return response.HandlerPostPlanetsResponse{
				PlanetName: planet.Name,
				Climate:    planet.Climate,
				Terrain:    planet.Terrain,
				Films:      len(planet.Films),
			}
		}
	}

	return response.HandlerPostPlanetsResponse{}
}
