package controllersv1

import (
	"net/http"

	"github.com/MohammadMobasher/resturan-backend/models"
	mongoRepositories "github.com/MohammadMobasher/resturan-backend/repositories/mongo_repository"

	"github.com/gin-gonic/gin"
)

// @Summary create a food group
// @Description create a food group
// @Tags food group
// @Accept */*
// @Produce json
// @Success 200
// @Router /v1/foodgroup [post]
func CreateFoodGroup(c *gin.Context) {
	var foodGroup models.FoodGroup
	err := c.ShouldBindJSON(&foodGroup)
	if err != nil {
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}

	foodGRoupRepository := mongoRepositories.NewFoodGroupRepository()
	reuslt, err := foodGRoupRepository.Insert(foodGroup)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, reuslt)
}

// @Summary delete a food group
// @Description delete a food group
// @Tags food group
// @Accept */*
// @Produce json
// @Success 200
// @Router /v1/foodgroup/:foodgroupId [delete]
func DeleteFoodGroup(c *gin.Context) {
	foodGroupId := c.Param("foodgroupId")
	foodGRoupRepository := mongoRepositories.NewFoodGroupRepository()
	result, err := foodGRoupRepository.Delete(foodGroupId)
	if err != nil && result {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "The foodGroup removed successfully"})
}

// @Summary Get all food group
// @Description Get all food group
// @Tags food group
// @Accept */*
// @Produce json
// @Success 200
// @Router /v1/foodgroup [Get]
func GetFoodGroups(c *gin.Context) {
	foodGRoupRepository := mongoRepositories.NewFoodGroupRepository()
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
	foodGRoupRepository := mongoRepositories.NewFoodGroupRepository()
	users, err := foodGRoupRepository.Update(foodGroup)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func GetFoodGroup(c *gin.Context) {
	userId := c.Param("foodgroupId")
	foodGRoupRepository := mongoRepositories.NewFoodGroupRepository()
	result, err := foodGRoupRepository.GetItem(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, result)
}
