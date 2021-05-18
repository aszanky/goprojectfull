package models

type Store struct {
	WPCode   string `json:"wp_code"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Location string `json:"location"`
}
