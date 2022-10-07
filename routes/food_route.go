package routes

import (
	controllersv1 "github.com/MohammadMobasher/resturan-backend/controllers/v1"
	"github.com/gin-gonic/gin"
)

func FoodRouteV1(r *gin.RouterGroup) {
	// r.GET("/foodgroup", controllers.GetFoodGroups)
	r.POST("/food", controllersv1.CreateFood)
	// r.PUT("/foodgroup", controllers.UpdateFoodGroup)
	// r.GET("/foodgroup/:foodgroupId", controllers.GetFoodGroup)
	// r.DELETE("/foodgroup/:foodgroupId", controllers.DeleteFoodGroup)
}
