package routes

import (
	usercontroller "github.com/MohammadMobasher/resturan-backend/controllers/user"
	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	r.GET("/users", usercontroller.GetUsers)
	r.POST("/user", usercontroller.CreateUser)
	r.PUT("/user", usercontroller.UpdateUsers)
	r.GET("/user/:userId", usercontroller.GetUser)
	r.DELETE("/user/:userId", usercontroller.DeleteUser)
}
