package routes

import (
	"github.com/gin-gonic/gin"
	"rms/database/middelware"
	"rms/handler"
)

func ServerRoutes(r1 *gin.Engine) {

	/*Login goes not need authentication , as at this time we just validate and assign token */
	r1.POST("/login", handler.Login)

	/*Admin level check*/
	userRouter := r1.Group("/admin-level")
	userRouter.Use(middelware.AdminMiddleware())
	{
		userRouter.POST("/create-user", handler.User)
		userRouter.GET("/all-users", handler.AllUsers)
		userRouter.GET("/all-SubAdmin", handler.AllSubAdmin)
	}

	/*SubAdmin level check*/
	userRouter = r1.Group("/sub-admin-level")
	userRouter.Use(middelware.UserMiddleware())
	{
		userRouter.POST("/customer", handler.User)
		userRouter.POST("/add-restaurant", handler.AddRestaurant)
		userRouter.POST("/new-dish", handler.AddDishes)
		userRouter.GET("/all-customers", handler.AllCustomers)
	}
	/* Customer level check*/
	userRouter = r1.Group("/user-level")
	userRouter.Use(middelware.AccessMiddleware())
	{
		userRouter.GET("/dish-by-restaurant", handler.Dish)
		userRouter.GET("/restaurants", handler.AllRestaurants)
		userRouter.GET("/distance", handler.Distance)
	}

	//r1.Run(":8080")
}
