package controllers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/Kunal-Patro/NoteTakingApp/dto"
	"github.com/Kunal-Patro/NoteTakingApp/initializers"
	"github.com/Kunal-Patro/NoteTakingApp/models"
	"github.com/Kunal-Patro/NoteTakingApp/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	var notebook models.Notebook
	initializers.DB.Find(&notebook, "id = ? AND user_id = ?", notebookID, user.(models.User).ID)

	if notebook.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot find the notebook.",
		})
		return
	}

	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to fetch page.",
		})
		return
	}

	var notesCount int64
	if err := initializers.DB.Table("notes").Count(&notesCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to count notes.",
		})
		return
	}

	pageCount := int(math.Ceil(float64(notesCount) / float64(NOTES_PER_PAGE)))

	if pageCount == 0 {
		pageCount = 1
	}

	if page < 1 || page > pageCount {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid page",
		})
		return
	}

	offset := (page - 1) * NOTES_PER_PAGE

	var notes []models.Note
	result := initializers.DB.Limit(NOTES_PER_PAGE).Offset(offset).
		Find(&notes, "notebook_id = ?", notebookID)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch notes.",
		})
		return
	}

	var prevPage, nextPage string

	if page > 1 {
		prevPage = fmt.Sprintf("%d", page-1)
	}
	if page < pageCount {
		nextPage = fmt.Sprintf("%d", page+1)
	}

	pages := make([]int, pageCount)

	for i := range pages {
		pages[i] = i + 1
	}

	c.JSON(http.StatusOK, gin.H{
		"notes":      notes,
		"next_page":  nextPage,
		"prev_page":  prevPage,
		"page_count": pageCount,
		"page":       page,
		"pages":      pages,
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
