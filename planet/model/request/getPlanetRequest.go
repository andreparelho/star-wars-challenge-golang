package request

type GetPlanetRequest struct {
	Name    string `json: "name"`
	Climate string `json: "climate"`
	Terrain string `json: "terrain"`
}
