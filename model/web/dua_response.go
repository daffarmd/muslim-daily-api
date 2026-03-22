package web

type DuaResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Arabic      string `json:"arabic"`
	Latin       string `json:"latin"`
	Translation string `json:"translation"`
	Category    string `json:"category"`
}
