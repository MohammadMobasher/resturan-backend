package routes

import (
	controllersv1 "github.com/MohammadMobasher/resturan-backend/controllers/v1"
	controllersv2 "github.com/MohammadMobasher/resturan-backend/controllers/v2"
	"github.com/gin-gonic/gin"
)

func UserRouteV1(r *gin.RouterGroup) {
	r.GET("/users", controllersv1.GetUsers)
	r.POST("/user", controllersv1.CreateUser)
	r.PUT("/user", controllersv1.UpdateUsers)
	r.GET("/user/:userId", controllersv1.GetUser)
	r.DELETE("/user/:userId", controllersv1.DeleteUser)
}

func UserRouteV2(r *gin.RouterGroup) {
	r.GET("/users", controllersv2.GetUsers)
	r.POST("/user", controllersv2.CreateUser)
	r.PUT("/user", controllersv2.UpdateUsers)
	r.GET("/user/:userId", controllersv2.GetUser)
	r.DELETE("/user/:userId", controllersv2.DeleteUser)
}
