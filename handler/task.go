package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rms/database"
	"rms/database/dbHelper"
	"rms/models"
)

func Dish(c *gin.Context) {

	var restaurantByDish []models.RestaurantByDishes
	restaurantByDish, err := dbHelper.RestaurantsByDish(database.Todo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"restaurants": restaurantByDish,
	})
	return

}
func AddDishes(c *gin.Context) {
	adminIdInterface, flag := c.Get("adminId")
	if !flag {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error to get id",
		})
		return
	}

	adminId, ok := adminIdInterface.(int)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "adminId is not of type int",
		})
		return
	}
	var dishes models.Dishes
	if err := c.BindJSON(&dishes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := dbHelper.AddDish(database.Todo, dishes, adminId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"restaurants": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Dish added successfully",
	})
	return
}
func AddRestaurant(c *gin.Context) {
	adminIdInterface, flag := c.Get("adminId")
	if !flag {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error to get id",
		})
		return
	}

	adminId, ok := adminIdInterface.(int)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "adminId is not of type int",
		})
		return
	}
	var restaurants models.Restaurants
	if err := c.BindJSON(&restaurants); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err := dbHelper.AddRestaurant(database.Todo, restaurants, adminId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Restaurant Added successfully",
	})
	return
}
func AllRestaurants(c *gin.Context) {

	var allRestaurants []models.Restaurants
	allRestaurants, err := dbHelper.AllRestaurants(database.Todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": allRestaurants,
	})
	return
}

//todo: fix status codes
func AllUsers(c *gin.Context) {
	var allInfo, err = dbHelper.AllUsers(database.Todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": allInfo,
	})
	return
}
func AllSubAdmin(c *gin.Context) {
	var AllInfo, err = dbHelper.AllSubAdmin(database.Todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": AllInfo,
	})
	return
}
func AllCustomers(c *gin.Context) {
	adminIdInterface, flag := c.Get("adminId")
	if !flag {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error to get id",
		})
		return
	}

	adminId, ok := adminIdInterface.(int)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "adminId is not of type int",
		})
		return
	}
	AllInfo, err := dbHelper.AllCustomers(database.Todo, adminId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": AllInfo,
	})
	return
}
func Distance(c *gin.Context) {
	var Coordinates models.Coordinates
	if err := c.BindJSON(&Coordinates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	distance, err := dbHelper.Distance(database.Todo, Coordinates.AddressId, Coordinates.Name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": distance,
	})

}
