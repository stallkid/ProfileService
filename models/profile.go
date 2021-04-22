package models

type Profile struct {
	Componenttitles []string `json:"componenttitles"`
	Persondata      [][]struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
		Data  string `json:"data"`
	} `json:"personData"`
}
