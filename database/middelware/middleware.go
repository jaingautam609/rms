package middelware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rms/database"
	"rms/database/authentication"
	"rms/database/dbHelper"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		adminId, err := ValidateToken(token)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		var flag bool
		flag, err = authentication.ValidateAdmin(database.Todo, adminId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"restaurants": err.Error(),
			})
			return
		}
		if flag == false {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "No access",
			})
			return
		}
		c.Set("adminId", adminId)
		c.Next()
	}
}
func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		adminId, err := ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		var flag bool
		flag, err = dbHelper.ValidateUser(database.Todo, adminId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"restaurants": err.Error(),
			})
			return
		}
		if flag == false {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "No access",
			})
			return
		}
		c.Next()
	}
}
func AccessMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		_, err := ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.Next()
	}
}
