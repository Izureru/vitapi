package main

type Meal struct {
	Name  string `json:"name"`
	Id    int    `json:"id"`
	Image string `json:"image"`
}

type Meals []Meal
