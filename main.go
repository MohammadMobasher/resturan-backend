package main

import (
	"net/http"

	"github.com/MohammadMobasher/resturan-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	createServer()

}

func createServer() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Homepage")
	})
	routes.UserRoute(r)
	r.Run()
}

// conf := config.GetConfig()
// ctx := context.TODO()

// db := database.ConnectDB(ctx, conf)

// if db != nil {
// 	fmt.Println("sucessfully")
// }
// user := models.User{
// 	Name:     "mohammad",
// 	UserName: "m.mobasher.z",
// }

// result, err := db.Collection("user").InsertOne(ctx, user)
// if err != nil {
// 	log.Println(err)
// 	panic(err)
// }

// fmt.Println(result)

// log.Println("Hello world!")
