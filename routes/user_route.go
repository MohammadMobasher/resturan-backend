package routes

import (
	controllersv1 "github.com/MohammadMobasher/resturan-backend/controllers/v1"
	"github.com/gin-gonic/gin"
)

func UserRouteV1(r *gin.RouterGroup) {
	r.GET("/users", controllersv1.GetUsers)
	r.POST("/user", controllersv1.CreateUser)
	r.PUT("/user", controllersv1.UpdateUsers)
	r.GET("/user/:userId", controllersv1.GetUser)
	r.DELETE("/user/:userId", controllersv1.DeleteUser)
}
