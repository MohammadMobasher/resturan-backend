package routes

import (
	"github.com/MohammadMobasher/resturan-backend/controllers"
	"github.com/gin-gonic/gin"
)

func FoodGroupRoute(r *gin.Engine) {
	r.GET("/foodgroup", controllers.GetFoodGroups)
	r.POST("/foodgroup", controllers.CreateFoodGroup)
	r.PUT("/foodgroup", controllers.UpdateFoodGroup)
	r.GET("/foodgroup/:foodgroupId", controllers.GetFoodGroup)
	r.DELETE("/foodgroup/:foodgroupId", controllers.DeleteFoodGroup)
}
