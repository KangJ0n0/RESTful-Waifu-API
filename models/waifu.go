package models

type Waifu struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Height float64 `json:"height"`
	Weight float64 `json:"weight"`
	Series string  `json:"series"`
	Pict   string  `json:"pict"`
}