package controllersv1

import (
	"net/http"

	"github.com/MohammadMobasher/resturan-backend/common"
	"github.com/MohammadMobasher/resturan-backend/models"
	mongoRepositories "github.com/MohammadMobasher/resturan-backend/repositories/mongo_repository"
	"github.com/gin-gonic/gin"
)

func CreateFood(c *gin.Context) {

	var food models.Food

	err := c.Bind(&food)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message1": err.Error()})
		return
	}
	imageAddress, err := common.UploadFile(c)
	food.ImageAddress = imageAddress

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message2": "1"})
		return
	}

	foodRepository := mongoRepositories.NewFoodRepository()
	reuslt, err := foodRepository.Insert(food)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, reuslt)
}
