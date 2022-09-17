package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	ErrorInternalError = errors.New("woops! something went wrong :(")
)

func ValidationErrorToText(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", e.Field())
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s", e.Field(), e.Param())
	case "min":
		return fmt.Sprintf("%s must be longer than %s", e.Field(), e.Param())
	case "email":
		return fmt.Sprintf("Invalid email format")
	case "len":
		return fmt.Sprintf("%s must be %s characters long", e.Field(), e.Param())
	}
	return fmt.Sprintf("%s is not valid", e.Field())
}

func ValidationErrors(c *gin.Context) {
	c.Next()
	if c.Errors != nil {
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				// 	switch e.Type {
				// 	case gin.ErrorTypePublic:
				// 		if !c.Writer.Written() {
				// 			c.JSON(c.Writer.Status(), gin.H{"Error": e.Error()})
				// 		}
				// 	case gin.ErrorTypeBind:
				// 		// log.Fatal("2")

				// 		var ve validator.ValidationErrors
				// 		list := make(map[string]interface{})
				// 		// log.Fatal("3")

				// 		if errors.As(e, &ve) {
				// 			// log.Fatal("4")

				// 			for _, field := range ve {
				// 				list[field.Field()] = field.Tag() //ValidationErrorToText(err)
				// 			}
				// 		}
				// 		// // Make sure we maintain the preset response status
				// 		status := http.StatusBadRequest
				// 		// log.Fatal("5")

				// 		if c.Writer.Status() != http.StatusOK {
				// 			status = c.Writer.Status()
				// 		}
				// 		c.IndentedJSON(status, list)
				// 	default:
				// 		// Log all other errors
				// 		// rollbar.RequestError(rollbar.ERR, c.Request, e.Err)
				// 		// if logger != nil {
				// 		// 	logger.Error(e.Err)
				// 		// }

				// 	}
				// }
				var ve validator.ValidationErrors
				list := make(map[string]interface{})

				if errors.As(e, &ve) {

					for _, field := range ve {
						list[field.Field()] = ValidationErrorToText(field)
					}
					status := http.StatusBadRequest
					if c.Writer.Status() != http.StatusOK {
						status = c.Writer.Status()
					}
					c.IndentedJSON(status, list)
				}

			}
		}

		if !c.Writer.Written() {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": ErrorInternalError.Error()})
		}
	}

}
