package routes

import (
	"github.com/gin-gonic/gin"
	endpoints "github.com/sanyogpatel-tecblic/CRUD_GORM/controller/EndPoints"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.POST("/users", endpoints.CreateUserHandler(db))
	router.GET("/users", endpoints.GetUserHandler(db))
	router.GET("/users/:id", endpoints.GetUserHandlerById(db))
	router.PUT("/users/:id", endpoints.UpdateUserHandler(db))
	router.DELETE("/users/:id", endpoints.DeleteUserHandler(db))
	return router
}
