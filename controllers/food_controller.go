package controllers

import (
	"net/http"

	"github.com/MohammadMobasher/resturan-backend/models"
	"github.com/MohammadMobasher/resturan-backend/repositories"
	"github.com/gin-gonic/gin"
)

func CreateFood(c *gin.Context) {

	var food models.Food

	err := c.Bind(&food)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message1": err.Error()})
		return
	}
	imageAddress, err := uploadFile(c)
	food.ImageAddress = imageAddress

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message2": "1"})
		return
	}

	foodRepository := repositories.NewFoodRepository()
	reuslt, err := foodRepository.Insert(food)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, reuslt)
}
