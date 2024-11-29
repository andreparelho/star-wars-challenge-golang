package response

type Result struct {
	Name       string   `json:"name"`
	Climate    string   `json:"climate"`
	Gravity    string   `json:"gravity"`
	Terrain    string   `json:"terrain"`
	Population string   `json:"population"`
	Films      []string `json:"films"`
}
