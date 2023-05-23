package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/sanyogpatel-tecblic/CRUD_GORM/controller/models"

	"gorm.io/gorm"
)

func CreateUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		//c.ShouldBindJSON to bind the JSON payload to the user variable

		if err := c.ShouldBindJSON(&user); err != nil {

		}
	}
}
