package main

import (
	"net/http"

	"github.com/MohammadMobasher/resturan-backend/routes"
	"github.com/MohammadMobasher/resturan-backend/routes/middleware"

	"github.com/gin-gonic/gin"

	_ "github.com/MohammadMobasher/resturan-backend/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin Swagger Example API
// @version 1.0
// @description resturan
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	// conf := config.GetConfig()
	// database.ConnectMySqlDB(conf)
	// foodGroup := models.FoodGroupMySql{
	// 	Name: "mohammad",
	// }
	// r := mysql_database.NewFoodGroupMySqlRepository()
	// r.Insert(foodGroup)

	createServer()

}

func createServer() {
	r := gin.Default()
	r.MaxMultipartMemory = 4 << 20 // 4 MiB
	r.Use(middleware.ValidationErrors)
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Homepage")
	})
	r.Static("/Content", "./uploaded_file")

	var rGroup = routes.ConfigAuth(r)

	rGroupV1 := rGroup.Group("/v1")

	routes.UserRouteV1(rGroupV1)
	routes.FoodRouteV1(rGroupV1)
	routes.FoodGroupRouteV1(rGroupV1)

	rGroupV2 := rGroup.Group("/v2")

	routes.FoodGroupRouteV2(rGroupV2)
	routes.FoodRouteV2(rGroupV2)
	routes.UserRouteV2(rGroupV2)
	routes.ResturanRouteV2(rGroupV2)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Run()
}
