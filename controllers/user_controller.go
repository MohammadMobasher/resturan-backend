package controllers

import (
	"net/http"

	"github.com/MohammadMobasher/resturan-backend/models"
	"github.com/MohammadMobasher/resturan-backend/repositories"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"mmm1": err})
		// log.Println("KKKKKKKKK%T", err.Error())
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}

	userRepository := repositories.NewUserRepository()
	reuslt, err := userRepository.Insert(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, reuslt)
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("userId")
	userRepository := repositories.NewUserRepository()
	result, err := userRepository.Delete(userId)
	if err != nil && result {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "The user removed successfully"})
}

func GetUsers(c *gin.Context) {
	userRepository := repositories.NewUserRepository()
	users, err := userRepository.GetAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, users)
}

func UpdateUsers(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	userRepository := repositories.NewUserRepository()
	users, err := userRepository.Update(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	userId := c.Param("userId")
	userRepository := repositories.NewUserRepository()
	result, err := userRepository.GetItem(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, result)
}
