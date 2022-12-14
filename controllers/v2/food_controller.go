package controllersv2

import (
	"log"
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
	// foodImages := models.FoodImage{}

	err := c.Bind(&food)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message1": err.Error()})
		return
	}
	imageAddresses, err := common.UploadFiles(c)
	// food.ImageAddress = &imageAddress

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message2": "1"})
		return
	}

	foodRepository := mysqlRepositories.NewFoodMySqlRepository()
	reuslt, err := foodRepository.Insert(food, imageAddresses)

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
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "The food removed successfully"})
}

// @Summary Get all food
// @Description Get all food
// @Tags food
// @Accept */*
// @Produce json
// @Param        page  query   integer false  "page"
// @Param        pagecount    query     integer    false  "pagecount"
// @Success 200
// @Router /v2/food [Get]
func GetFoods(c *gin.Context) {
	pagination := models.Pagination{}
	err := c.BindQuery(&pagination)

	foodGRoupRepository := mysqlRepositories.NewFoodMySqlRepository()
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

// @Summary Get a food
// @Description Get a food
// @Tags food
// @Accept */*
// @Produce json
// @Success 200
// @Router /v2/food/:foodId [Get]
func GetFood(c *gin.Context) {
	foodId, err := strconv.ParseInt(c.Param("foodId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	foodGRoupRepository := mysqlRepositories.NewFoodMySqlRepository()
	result, err := foodGRoupRepository.GetItem(foodId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

// @Summary Create food Comment
// @Description Create food Comment
// @Tags food
// @Accept */*
// @Produce json
// @Success 200
// @Router /v2/food/comment/:foodId [Post]
func CreateFoodComment(c *gin.Context) {
	var foodComment models.FoodCommentMySql

	foodId, err := strconv.ParseInt(c.Param("foodId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	log.Println(foodId)
	foodComment.FoodId = foodId
	err = c.Bind(&foodComment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	foodRepository := mysqlRepositories.NewFoodMySqlRepository()
	_, err = foodRepository.CreateComment(foodComment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

// @Summary Get Comments
// @Description Get Comments
// @Tags food
// @Accept */*
// @Produce json
// @Param        page  query   integer false  "page"
// @Param        pagecount    query     integer    false  "pagecount"
// @Success 200
// @Router /v2/food/comment/:foodId [Get]
func GetComments(c *gin.Context) {
	foodId, err := strconv.ParseInt(c.Param("foodId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	pagination := models.Pagination{}
	err = c.BindQuery(&pagination)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	foodRepository := mysqlRepositories.NewFoodMySqlRepository()
	items, count, err := foodRepository.GetComments(foodId, pagination.Page*pagination.PageCount, pagination.PageCount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK,
		models.PagedResult{
			TotalCount: count,
			Items:      items,
		})
}
