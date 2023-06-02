package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanyogpatel-tecblic/CRUD_GORM/controller/models"

	"gorm.io/gorm"
)

func CreateUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		//c.ShouldBindJSON to bind the JSON payload to the user variable

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//Create() automatically generates an SQL INSERT statement based on the provided struct and inserts the corresponding record into the database table.
		result := db.Create(&user)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func GetUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a channel to receive the result from the goroutine
		resultChan := make(chan []models.User)

		go func() {
			var users []models.User

			// Perform the database operation
			result := db.Find(&users)
			if result.Error != nil {
				// Send the error to the channel
				resultChan <- nil
				return
			}

			// Send the result to the channel
			resultChan <- users
		}()

		// Wait for the goroutine to complete and receive the result from the channel
		users := <-resultChan

		// Check if the result is nil, indicating an error occurred
		if users == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

func GetUserHandlerById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var users models.User

		//Fisrt() function is used to retrieve the first record that matches the specified conditions from the database
		result := db.First(&users, id)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func UpdateUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var users models.User

		//First() function is used to retrieve the first record that matches the specified conditions from the database
		result := db.First(&users, id)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		if err := c.ShouldBindJSON(&users); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//Save() function is used to save or update a record in the database.

		result = db.Save(&users)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func DeleteUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var users models.User

		//First() function is used to retrieve the first record that matches the specified conditions from the database
		result := db.First(&users, id)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		//Save() function is used to save or update a record in the database.
		result = db.Delete(&users)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User Deleted Successfully"})
	}
}
