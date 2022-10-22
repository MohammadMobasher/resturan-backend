package controllersv2

import (
	"net/http"

	"github.com/MohammadMobasher/resturan-backend/models"
	mysqlRepositories "github.com/MohammadMobasher/resturan-backend/repositories/mysql_repository"

	"github.com/gin-gonic/gin"
)

// @Summary create a resturan
// @Description create a resturan
// @Tags resturan
// @Accept */*
// @Produce json
// @Success 200
// @Router /v2/resturan [post]
func CreateResturan(c *gin.Context) {

	var resturan models.ResturanMySql

	err := c.Bind(&resturan)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message1": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message2": "1"})
		return
	}

	foodRepository := mysqlRepositories.NewResturanMySqlRepository()
	reuslt, err := foodRepository.Insert(resturan)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, reuslt)
}
