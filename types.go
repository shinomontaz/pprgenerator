package generator

import "time"

type LatLong struct {
	Lat  float64
	Long float64
}

type Job struct {
	Id          int
	Name        string
	Difficulty  float64 // человеко-часов на выполнение этой работы
	Days        []time.Time
	FkEquipment *Equipment
	IsBreaking  bool
}

type Tower struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Weight int    `json:"weight"` // охват населения
	Coords LatLong
}

type Equipment struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Weight   int     `json:"weight"`   // охват населения
	Majority float64 `json:"majority"` // важность
	FkTower  *Tower
}

type Interrupt struct {
	Day    time.Time
	Length float64 // длительность остановки вещания в часах
}

type Stop struct {
	Day         time.Time
	Length      float64 // длительность остановки вещания в часах
	FkEquipment *Equipment
}
