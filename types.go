package generator

import "time"

type Job struct {
	Id         int
	Name       string
	Difficulty float64 // человеко-часов на выполнение этой работы
	Days       []time.Time
	pTower     *Tower
	IsBreaking bool
}

type Tower struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Weight int     `json:"weight"` // охват населения
	x      float64 `json:"x"`
	y      float64 `json:"y"`
}
