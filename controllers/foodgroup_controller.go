package controllers

import (
	"net/http"

	"github.com/MohammadMobasher/resturan-backend/models"
	"github.com/MohammadMobasher/resturan-backend/repositories"
	"github.com/gin-gonic/gin"
)

func CreateFoodGroup(c *gin.Context) {
	var foodGroup models.FoodGroup
	err := c.ShouldBindJSON(&foodGroup)
	if err != nil {
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}

	foodGRoupRepository := repositories.NewFoodGroupRepository()
	reuslt, err := foodGRoupRepository.Insert(foodGroup)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, reuslt)
}

func DeleteFoodGroup(c *gin.Context) {
	userId := c.Param("foodgroupId")
	foodGRoupRepository := repositories.NewFoodGroupRepository()
	result, err := foodGRoupRepository.Delete(userId)
	if err != nil && result {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "The foodGroup removed successfully"})
}

func GetFoodGroups(c *gin.Context) {
	foodGRoupRepository := repositories.NewFoodGroupRepository()
	users, err := foodGRoupRepository.GetAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, users)
}

func UpdateFoodGroup(c *gin.Context) {
	foodGroup := models.FoodGroup{}
	err := c.ShouldBindJSON(&foodGroup)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	foodGRoupRepository := repositories.NewFoodGroupRepository()
	users, err := foodGRoupRepository.Update(foodGroup)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func GetFoodGroup(c *gin.Context) {
	userId := c.Param("foodgroupId")
	foodGRoupRepository := repositories.NewFoodGroupRepository()
	result, err := foodGRoupRepository.GetItem(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, result)
}
