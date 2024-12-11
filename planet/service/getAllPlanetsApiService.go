package service

import (
	"encoding/json"
	"net/http"

	"github.com/andreparelho/star-wars-challenge-golang/planet/common"
	"github.com/andreparelho/star-wars-challenge-golang/planet/model/response"
	"github.com/gin-gonic/gin"
)

const QUERY_PARAM string = "?page="
const PAGE string = "page"

type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

type RealHTTPClient struct{}

func (realHttpClient *RealHTTPClient) Get(url string) (*http.Response, error) {
	return http.Get(url)
}

type GetAllPlanetsService struct {
	Client HTTPClient
}

func (service *GetAllPlanetsService) GetAllPlanets(context *gin.Context) (response.HandlerGetPlanetsResponse, error) {
	var page string = context.Query(PAGE)

	requestApi, error := service.Client.Get(common.API_URL + QUERY_PARAM + page)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			common.ERROR_KEY: common.CALL_API_ERROR_MESSAGE,
		})
		return response.HandlerGetPlanetsResponse{}, error
	}

	defer requestApi.Body.Close()

	var getPlanetsResponse = response.Data{}
	error = json.NewDecoder(requestApi.Body).Decode(&getPlanetsResponse)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			common.ERROR_KEY: common.JSON_PROCESS_ERROR_MESSAGE,
		})
		return response.HandlerGetPlanetsResponse{}, error
	}

	var responseHandler = response.HandlerGetPlanetsResponse{
		Planets: getPlanetsResponse.Result,
	}

	return responseHandler, nil
}
