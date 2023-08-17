package controllers

import (
	"net/http"

	"github.com/Kunal-Patro/NoteTakingApp/initializers"
	"github.com/Kunal-Patro/NoteTakingApp/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateNotebook(c *gin.Context) {
	var body struct {
		Name string `json:"name"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	user, _ := c.Get("user")

	// if !ok {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Failed to fetch user",
	// 	})
	// 	return
	// }

	notebook := models.Notebook{
		Name: body.Name,
		User: user.(models.User),
	}

	result := initializers.DB.Create(&notebook)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create notebook",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": notebook,
	})
}

func GetNotebooks(c *gin.Context) {
	user, _ := c.Get("user")

	var notebooks []models.Notebook
	initializers.DB.Find(&notebooks, "user_id = ?", user.(models.User).ID)

	// for i := range notebooks {
	// 	notebooks[i].User = user.(models.User)
	// }

	if notebooks == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Cannot find any notebooks created by the user.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"notebooks": notebooks,
	})
}

func GetNotebook(c *gin.Context) {
	notebookID := c.Param("notebook_id")

	var notebook models.Notebook
	initializers.DB.Find(&notebook, "id = ?", notebookID)

	if notebook.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot find the requested notebook.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"notebook": notebook,
	})
}

func UpdateNotebook(c *gin.Context) {
	var body struct {
		Name string `json:"name"`
	}

	user, _ := c.Get("user")

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	notebookID := c.Param("notebook_id")

	var notebook models.Notebook
	initializers.DB.Find(&notebook, "id = ?", notebookID)

	if notebook.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot find the requested notebook.",
		})
		return
	}

	if notebook.UserID != user.(models.User).ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Operation",
		})
		return
	}

	result := initializers.DB.Model(&notebook).Updates(models.Notebook{Name: body.Name})

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update notebook.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Record Updated",
	})

}

func DeleteNotebook(c *gin.Context) {
	user, _ := c.Get("user")
	notebookID := c.Param("notebook_id")

	var notebook models.Notebook
	initializers.DB.Find(&notebook, "id = ?", notebookID)

	if notebook.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot find the requested notebook.",
		})
		return
	}

	if notebook.UserID != user.(models.User).ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid operation",
		})
		return
	}

	result := initializers.DB.Delete(&notebook)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete notebook.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Notebook deleted.",
	})
}
