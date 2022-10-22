package routes

import (
	controllersv1 "github.com/MohammadMobasher/resturan-backend/controllers/v1"
	controllersv2 "github.com/MohammadMobasher/resturan-backend/controllers/v2"

	"github.com/gin-gonic/gin"
)

func FoodRouteV1(r *gin.RouterGroup) {
	// r.GET("/foodgroup", controllers.GetFoodGroups)
	r.POST("/food", controllersv1.CreateFood)
	// r.PUT("/foodgroup", controllers.UpdateFoodGroup)
	// r.GET("/foodgroup/:foodgroupId", controllers.GetFoodGroup)
	// r.DELETE("/foodgroup/:foodgroupId", controllers.DeleteFoodGroup)
}

func FoodRouteV2(r *gin.RouterGroup) {
	r.GET("/food", controllersv2.GetFoods)
	r.POST("/food", controllersv2.CreateFood)
	// r.PUT("/foodgroup", controllers.UpdateFoodGroup)
	// r.GET("/foodgroup/:foodgroupId", controllers.GetFoodGroup)
	r.DELETE("/food/:id", controllersv2.DeleteFood)
}
