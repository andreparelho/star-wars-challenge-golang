package response

type HandlerPostPlanetsResponse struct {
	PlanetName string `json:"planet"`
	Climate    string `json:"climate"`
	Terrain    string `json:"terrain"`
	Films      int    `json:"films"`
}
