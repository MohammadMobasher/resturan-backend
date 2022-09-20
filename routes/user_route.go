package routes

import (
	"github.com/MohammadMobasher/resturan-backend/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	r.GET("/users", controllers.GetUsers)
	r.POST("/user", controllers.CreateUser)
	r.PUT("/user", controllers.UpdateUsers)
	r.GET("/user/:userId", controllers.GetUser)
	r.DELETE("/user/:userId", controllers.DeleteUser)
}
