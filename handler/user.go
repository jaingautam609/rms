package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rms/database"
	"rms/database/authentication"
	_ "rms/database/dbHelper"
	"rms/database/middelware"
	"rms/models"
)

func User(c *gin.Context) {
	//todo: this token should not be here
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
			"message": "cannot convert to int type",
		})
		return
	}

	var addUser models.AddUser
	if err := c.BindJSON(&addUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	tx, err := database.Todo.Begin() /*could be here as it is .Todo*/
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	userId, err := authentication.Create(database.Todo, adminId, addUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		_ = tx.Rollback()
		return
	}
	err = authentication.AddRole(database.Todo, userId, addUser.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		_ = tx.Rollback()
		return
	}
	err = authentication.AddAddress(database.Todo, userId, addUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		_ = tx.Rollback()
		return
	}
	err = tx.Commit()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Created Successful",
	})
	return
}

func Login(c *gin.Context) {
	var user models.Users
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	uId, err := authentication.Login(database.Todo, user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	signedToken, err := middelware.GenerateToken(uId, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": signedToken,
	})
	return
}
