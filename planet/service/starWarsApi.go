package service

import (
	"log"
	"net/http"
)

type CreatePlanetRequest struct {
	Name string `json: "name"`
}

const apiUrl = "https://swapi.dev/api/planets/"

func GetPlanets(createPlanetRequest CreatePlanetRequest) *http.Response {
	response, error := http.Get(apiUrl)
	if error != nil {
		log.Fatal("error request api", error)
	}

	return response
}
