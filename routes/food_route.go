package routes

import (
	"github.com/MohammadMobasher/resturan-backend/controllers"
	"github.com/gin-gonic/gin"
)

func FoodRoute(r *gin.RouterGroup) {
	// r.GET("/foodgroup", controllers.GetFoodGroups)
	r.POST("/food", controllers.CreateFood)
	// r.PUT("/foodgroup", controllers.UpdateFoodGroup)
	// r.GET("/foodgroup/:foodgroupId", controllers.GetFoodGroup)
	// r.DELETE("/foodgroup/:foodgroupId", controllers.DeleteFoodGroup)
}
