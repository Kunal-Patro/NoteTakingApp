package controllers

import (
	"fmt"
	"net/http"

	"github.com/Kunal-Patro/NoteTakingApp/initializers"
	"github.com/Kunal-Patro/NoteTakingApp/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SearchContent(c *gin.Context) {
	var body struct {
		Content string `json:"content"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot read body.",
		})
		return
	}

	user, _ := c.Get("user")
	body.Content = fmt.Sprintf("%v%v%v", "%", body.Content, "%")
	var notebooks []models.Notebook
	result := initializers.DB.Find(&notebooks, "user_id = ?", user.(models.User).ID)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to find notebooks",
		})
		return
	}

	var notebookIDs []uuid.UUID
	for _, notebook := range notebooks {
		notebookIDs = append(notebookIDs, notebook.ID)
	}

	result = initializers.DB.Where("name LIKE ? AND user_id = ?", body.Content, user.(models.User).ID).Find(&notebooks)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch notebook data",
		})
		return
	}
	var notes []models.Note
	result = initializers.DB.Where("(title LIKE ? OR description LIKE ?) AND notebook_id IN ?", body.Content, body.Content, notebookIDs).Find(&notes)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch notes data.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"notebooks": notebooks,
		"notes":     notes,
	})

}
