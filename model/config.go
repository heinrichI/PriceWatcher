package model

// Config Настройки
type Config struct {
	Item struct {
		Name	string `json:"name"`
		Urls string[] `json:"urls"`
	} `json:"input"`
	Percent       int    `json:"percent"`
	Items	[]Item `json:"items"`
}
