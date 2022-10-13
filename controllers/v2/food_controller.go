package controllersv2

import (
	"net/http"

	"github.com/MohammadMobasher/resturan-backend/common"
	"github.com/MohammadMobasher/resturan-backend/models"
	mysqlRepositories "github.com/MohammadMobasher/resturan-backend/repositories/mysql_repository"

	"github.com/gin-gonic/gin"
)

// @Summary create a food
// @Description create a food
// @Tags food
// @Accept */*
// @Produce json
// @Param        Name  query   string false  "food name"
// @Param        file  formData   file false  "food image"
// @Success 200
// @Router /v2/food [post]
func CreateFood(c *gin.Context) {

	var food models.FoodMySql

	err := c.Bind(&food)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message1": err.Error()})
		return
	}
	imageAddress, err := common.UploadFile(c)
	food.ImageAddress = &imageAddress

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message2": "1"})
		return
	}

	foodRepository := mysqlRepositories.NewFoodMySqlRepository()
	reuslt, err := foodRepository.Insert(food)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, reuslt)
}
