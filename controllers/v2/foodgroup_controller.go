package controllersv2

import (
	"log"
	"net/http"
	"strconv"

	"github.com/MohammadMobasher/resturan-backend/common"
	mysqlRepositories "github.com/MohammadMobasher/resturan-backend/repositories/mysql_repository"

	"github.com/MohammadMobasher/resturan-backend/models"
	"github.com/gin-gonic/gin"
)

// @Summary create a food group
// @Description create a food group
// @Tags food group
// @Accept */*
// @Produce json
// @Param        Name  query   string false  "food group name"
// @Param        file  formData   file false  "food group image"
// @Success 200
// @Router /v2/foodgroup [post]
func CreateFoodGroup(c *gin.Context) {
	var foodGroup models.FoodGroupMySql
	err := c.Bind(&foodGroup)
	if err != nil {
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}

	imageAddress, err := common.UploadFile(c)
	foodGroup.ImageAddress = &imageAddress

	foodGRoupRepository := mysqlRepositories.NewFoodGroupMySqlRepository()
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
// @Router /v2/foodgroup/:foodgroupId [delete]
func DeleteFoodGroup(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("foodgroupId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	foodGRoupRepository := mysqlRepositories.NewFoodGroupMySqlRepository()
	result, err := foodGRoupRepository.Delete(id)
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
// @Router /v2/foodgroup [Get]
func GetFoodGroups(c *gin.Context) {
	pagination := models.Pagination{}
	err := c.BindQuery(&pagination)

	log.Println(pagination.Page)
	log.Println(pagination.PageCount)

	foodGRoupRepository := mysqlRepositories.NewFoodGroupMySqlRepository()
	foodGroups, err := foodGRoupRepository.GetAll(pagination.Page*pagination.PageCount, pagination.PageCount)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, foodGroups)
}

// @Summary Get a food group
// @Description Get a food group
// @Tags food group
// @Accept */*
// @Produce json
// @Success 200
// @Router /v2/foodgroup/:foodgroupId [Get]
func GetFoodGroup(c *gin.Context) {
	foodgroupId, err := strconv.ParseInt(c.Param("foodgroupId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	foodGRoupRepository := mysqlRepositories.NewFoodGroupMySqlRepository()
	result, err := foodGRoupRepository.GetItem(foodgroupId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, result)
}

// @Summary Update a food group
// @Description Update a food group
// @Tags food group
// @Accept */*
// @Produce json
// @Param        Id    query     integer    false  "food group id"
// @Param        Name  query   string false  "food group name"
// @Param        file  formData   file false  "food group image"
// @Success 200
// @Router /v2/foodgroup [PUT]
func UpdateFoodGroup(c *gin.Context) {
	foodGroup := models.FoodGroupMySql{}
	err := c.Bind(&foodGroup)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	imageAddress, err := common.UploadFile(c)
	if imageAddress != "" {
		foodGroup.ImageAddress = &imageAddress
	}

	foodGRoupRepository := mysqlRepositories.NewFoodGroupMySqlRepository()
	foodGroups, err := foodGRoupRepository.Update(foodGroup)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, foodGroups)
}
