package controllersv2

import (
	"net/http"
	"strconv"

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

// @Summary Get all resturan
// @Description Get all resturan
// @Tags resturan
// @Accept */*
// @Produce json
// @Param        page  query   integer false  "page"
// @Param        pagecount    query     integer    false  "pagecount"
// @Success 200
// @Router /v2/resturan [Get]
func GetResturans(c *gin.Context) {
	pagination := models.Pagination{}
	err := c.BindQuery(&pagination)

	foodGRoupRepository := mysqlRepositories.NewResturanMySqlRepository()
	foodGroups, count, err := foodGRoupRepository.GetAll(pagination.Page*pagination.PageCount, pagination.PageCount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK,
		models.PagedResult{
			TotalCount: count,
			Items:      foodGroups,
		})
}

// @Summary delete a resturan
// @Description delete a resturan
// @Tags resturan
// @Accept */*
// @Produce json
// @Success 200
// @Router /v2/resturan/:resturanId [delete]
func DeleteResturan(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("resturanId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	foodGRoupRepository := mysqlRepositories.NewResturanMySqlRepository()
	result, err := foodGRoupRepository.Delete(id)
	if err != nil && result {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "The resturan removed successfully"})
}
