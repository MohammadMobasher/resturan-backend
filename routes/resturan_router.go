package routes

import (
	controllersv2 "github.com/MohammadMobasher/resturan-backend/controllers/v2"

	"github.com/gin-gonic/gin"
)

// func FoodRouteV1(r *gin.RouterGroup) {
// 	// r.GET("/foodgroup", controllers.GetFoodGroups)
// 	r.POST("/food", controllersv1.CreateFood)
// 	// r.PUT("/foodgroup", controllers.UpdateFoodGroup)
// 	// r.GET("/foodgroup/:foodgroupId", controllers.GetFoodGroup)
// 	// r.DELETE("/foodgroup/:foodgroupId", controllers.DeleteFoodGroup)
// }

func ResturanRouteV2(r *gin.RouterGroup) {
	r.GET("/resturan", controllersv2.GetResturans)
	r.POST("/resturan", controllersv2.CreateResturan)
	// r.PUT("/foodgroup", controllers.UpdateFoodGroup)
	// r.GET("/food/:foodId", controllersv2.GetFood)
	r.DELETE("/food/:id", controllersv2.DeleteResturan)
}
