package models

import "time"

type Restaurants struct {
	Id          int
	Name        string    `json:"name"`
	Longitude   float64   `json:"longitude"`
	Latitude    float64   `json:"latitude"`
	Address     string    `json:"address"`
	OpeningTime time.Time `json:"openingTime"`
	ClosingTime time.Time `json:"closingTime"`
	UsedId      int
}
type Users struct {
	Id        int    `json:"id"`
	Name      string `json:"userName"`
	Email     string `json:"userEmail"`
	Password  string `json:"password"`
	CreatedBy int    `json:"createdBy"`
}
type Dishes struct {
	Id           int
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	RestaurantId int       `json:"restaurantId"`
	StartServe   time.Time `json:"startServe"`
	EndServe     time.Time `json:"endServe"`
}
type AllInfo struct {
	RestaurantName string    `json:"restaurantName"`
	Address        string    `json:"address"`
	OpeningTime    time.Time `json:"openingTime"`
	ClosingTime    time.Time `json:"closingTime"`
	DishName       string    `json:"dishName"`
	Type           string    `json:"type"`
	StartServe     time.Time `json:"startServe"`
	EndServe       time.Time `json:"endServe"`
	UserName       string    `json:"userName"`
	UserEmail      string    `json:"userEmail"`
}
type RestaurantByDishes struct {
	RestaurantName string    `json:"restaurantName"`
	Address        string    `json:"address"`
	OpeningTime    time.Time `json:"openingTime"`
	ClosingTime    time.Time `json:"closingTime"`
	DishName       string    `json:"dishName"`
	Type           string    `json:"type"`
	StartServe     time.Time `json:"startServe"`
	EndServe       time.Time `json:"endServe"`
}
type CustomersBySubAdmin struct {
	RestaurantName string    `json:"restaurantName"`
	Address        float64   `json:"address"`
	OpeningTime    time.Time `json:"openingTime"`
	ClosingTime    time.Time `json:"closingTime"`
	UserName       string    `json:"userName"`
	UserEmail      string    `json:"userEmail"`
}
type Distance struct {
	ResLongi  float64
	ResLati   float64
	UserLongi float64
	UserLati  float64
}
type AddUser struct {
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Role      string  `json:"role"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Location  string  `json:"location"`
}
type Coordinates struct {
	Name      string `json:"name"`
	AddressId int    `json:"addressId"`
}
