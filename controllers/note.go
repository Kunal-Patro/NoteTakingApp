package controllers

import (
	"net/http"

	"github.com/Kunal-Patro/NoteTakingApp/dto"
	"github.com/Kunal-Patro/NoteTakingApp/models"
	"github.com/Kunal-Patro/NoteTakingApp/services"
	"github.com/gin-gonic/gin"
)

const NOTES_PER_PAGE = 3

func CreateNote(c *gin.Context) {
	var body dto.NoteDTO

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	user, _ := c.Get("user")
	notebookID := c.Param("notebook_id")

	res := services.CreateNewNote(&body, notebookID, user.(models.User))

	tag := "message"
	if res.Code != http.StatusOK {
		tag = "error"
	}

	c.JSON(res.Code, gin.H{
		tag: res.Body,
	})

}

func GetAllNotes(c *gin.Context) {
	user, _ := c.Get("user")
	notebookID := c.Param("notebook_id")
	pageStr := c.DefaultQuery("page", "1")

	res := services.FetchAllNotes(notebookID, pageStr, user.(models.User))

	tag := "notes"
	if res.Code != http.StatusOK {
		tag = "error"
	}

	c.JSON(res.Code, gin.H{
		tag:         res.Body,
		"page_data": res.Page, // page data will be {"page_count": 0,"curr_page": 0,"prev_page": "","next_page": "","pages": null} in case of error.
	})
}

func GetNote(c *gin.Context) {
	user, _ := c.Get("user")
	notebookID := c.Param("notebook_id")
	noteID := c.Param("note_id")

	res := services.FetchNote(noteID, notebookID, user.(models.User))

	tag := "notes"
	if res.Code != http.StatusOK {
		tag = "error"
	}

	c.JSON(res.Code, gin.H{
		tag: res.Body,
	})
}

func UpdateNote(c *gin.Context) {
	var body dto.NoteDTO

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body.",
		})
		return
	}

	user, _ := c.Get("user")
	notebookID := c.Param("notebook_id")
	noteID := c.Param("note_id")

	res := services.AlterNote(&body, noteID, notebookID, user.(models.User))

	tag := "message"

	if res.Code != http.StatusOK {
		tag = "error"
	}

	c.JSON(res.Code, gin.H{
		tag: res.Body,
	})
}

func DeleteNote(c *gin.Context) {
	user, _ := c.Get("user")
	notebookID := c.Param("notebook_id")
	noteID := c.Param("note_id")

	res := services.RemoveNote(noteID, notebookID, user.(models.User))

	tag := "message"

	if res.Code != http.StatusOK {
		tag = "error"
	}

	c.JSON(res.Code, gin.H{
		tag: res.Body,
	})
}
