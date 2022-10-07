package controllersv1

import (
	"net/http"

	"github.com/MohammadMobasher/resturan-backend/models"
	"github.com/MohammadMobasher/resturan-backend/repositories"
	"github.com/gin-gonic/gin"
)

// @Summary create a user
// @Description create a user
// @Tags user
// @Accept */*
// @Produce json
// @Success 200
// @Router /v1/user [post]
func CreateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}

	userRepository := repositories.NewUserRepository()
	reuslt, err := userRepository.Insert(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, reuslt)
}

// @Summary delete a user
// @Description delete a user
// @Tags user
// @Accept */*
// @Produce json
// @Success 200
// @Router /user/:userId [delete]
func DeleteUser(c *gin.Context) {
	userId := c.Param("userId")
	userRepository := repositories.NewUserRepository()
	result, err := userRepository.Delete(userId)
	if err != nil && result {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "The user removed successfully"})
}

// @Summary Get all users.
// @Description Get all users
// @Tags user
// @Accept */*
// @Produce json
// @Success 200
// @Router /users [get]
func GetUsers(c *gin.Context) {
	userRepository := repositories.NewUserRepository()
	users, err := userRepository.GetAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, users)
}

// @Summary put a user
// @Description put a user
// @Tags user
// @Accept */*
// @Produce json
// @Success 200
// @Router /user [put]
func UpdateUsers(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	userRepository := repositories.NewUserRepository()
	users, err := userRepository.Update(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

// @Summary get a user
// @Description get a user
// @Tags user
// @Accept */*
// @Produce json
// @Success 200
// @Router /user [get]
func GetUser(c *gin.Context) {
	userId := c.Param("userId")
	userRepository := repositories.NewUserRepository()
	result, err := userRepository.GetItem(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, result)
}
