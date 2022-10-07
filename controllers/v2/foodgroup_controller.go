package controllersv2

import (
	"net/http"

	mysqlRepositories "github.com/MohammadMobasher/resturan-backend/repositories/mysql_repository"

	"github.com/MohammadMobasher/resturan-backend/models"
	"github.com/gin-gonic/gin"
)

func CreateFoodGroup(c *gin.Context) {
	var foodGroup models.FoodGroupMySql
	err := c.ShouldBindJSON(&foodGroup)
	if err != nil {
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}

	foodGRoupRepository := mysqlRepositories.NewFoodGroupMySqlRepository()
	reuslt, err := foodGRoupRepository.Insert(foodGroup)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, reuslt)
}
