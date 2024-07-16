package model

type Response struct {
	Name   string
	Tribes []Tribe
	Page   int
	Prev   int
	Next   int
}
