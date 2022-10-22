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
	r.PUT("/resturan", controllersv2.UpdateResturan)
	r.GET("/resturan/:resturanId", controllersv2.GetResturan)
	r.DELETE("/food/:id", controllersv2.DeleteResturan)
}
