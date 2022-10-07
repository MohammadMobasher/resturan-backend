package routes

import (
	controllersv1 "github.com/MohammadMobasher/resturan-backend/controllers/v1"
	controllersv2 "github.com/MohammadMobasher/resturan-backend/controllers/v2"
	"github.com/gin-gonic/gin"
)

func FoodGroupRouteV1(r *gin.RouterGroup) {
	r.GET("/foodgroup", controllersv1.GetFoodGroups)
	r.POST("/foodgroup", controllersv1.CreateFoodGroup)
	r.PUT("/foodgroup", controllersv1.UpdateFoodGroup)
	r.GET("/foodgroup/:foodgroupId", controllersv1.GetFoodGroup)
	r.DELETE("/foodgroup/:foodgroupId", controllersv1.DeleteFoodGroup)
}

func FoodGroupRouteV2(r *gin.RouterGroup) {
	// r.GET("/foodgroup", controllersv1.GetFoodGroups)
	r.POST("/foodgroup", controllersv2.CreateFoodGroup)
	// r.PUT("/foodgroup", controllersv1.UpdateFoodGroup)
	// r.GET("/foodgroup/:foodgroupId", controllersv1.GetFoodGroup)
	r.DELETE("/foodgroup/:foodgroupId", controllersv2.DeleteFoodGroup)
}
