package controllersv2

import (
	"net/http"
	"strconv"

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
// @Param        FoodGroupId    query     integer    false  "food group id"
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

// @Summary delete a food
// @Description delete a food
// @Tags food
// @Accept */*
// @Produce json
// @Param        Name  query   string false  "food name"
// @Success 200
// @Router /v2/food/:id [delete]
func DeleteFood(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	foodRepository := mysqlRepositories.NewFoodMySqlRepository()
	result, err := foodRepository.Delete(id)
	if err != nil && result {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "The food removed successfully"})
}
