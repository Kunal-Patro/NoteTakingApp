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
)

const NOTEBOOK_PER_PAGE = 3

func CreateNotebook(c *gin.Context) {
	var body dto.NotebookDTO

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	user, _ := c.Get("user")

	res := services.CreateNotebook(&body, user.(models.User))

	tag := "message"
	if res.Code != http.StatusOK {
		tag = "error"
	}
	c.JSON(res.Code, gin.H{
		tag: res.Body,
	})
}

func GetAllNotebooks(c *gin.Context) {
	user, _ := c.Get("user")

	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to get page value.",
		})
		return
	}

	var notebookCount int64
	if err := initializers.DB.Table("notebooks").Count(&notebookCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to count notebooks",
		})
		return
	}

	pageCount := int(math.Ceil(float64(notebookCount) / float64(NOTEBOOK_PER_PAGE)))

	if pageCount == 0 {
		pageCount = 1
	}
	if page < 1 || page > pageCount {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Page",
		})
		return
	}

	offset := (page - 1) * NOTEBOOK_PER_PAGE

	var notebooks []models.Notebook
	result := initializers.DB.Limit(NOTEBOOK_PER_PAGE).Offset(offset).
		Find(&notebooks, "user_id = ?", user.(models.User).ID)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch data.",
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

	for i := 0; i < pageCount; i++ {
		pages[i] = i + 1
	}

	c.JSON(http.StatusOK, gin.H{
		"notebooks":  notebooks,
		"page_count": pageCount,
		"page":       page,
		"prev_page":  prevPage,
		"next_page":  nextPage,
		"pages":      pages,
	})
}

func GetNotebook(c *gin.Context) {
	user, _ := c.Get("user")
	notebookID := c.Param("notebook_id")

	res := services.FetchNotebook(notebookID, user.(models.User))

	tag := "notebook"
	if res.Code != http.StatusOK {
		tag = "error"
	}
	c.JSON(res.Code, gin.H{
		tag: res.Body,
	})
}

func UpdateNotebook(c *gin.Context) {
	var body dto.NotebookDTO

	user, _ := c.Get("user")
	notebookID := c.Param("notebook_id")

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	res := services.AlterNotebook(&body, notebookID, user.(models.User))

	tag := "message"
	if res.Code != http.StatusOK {
		tag = "error"
	}

	c.JSON(res.Code, gin.H{
		tag: res.Body,
	})

}

func DeleteNotebook(c *gin.Context) {
	user, _ := c.Get("user")
	notebookID := c.Param("notebook_id")

	res := services.RemoveNotebook(notebookID, user.(models.User))

	tag := "message"
	if res.Code != http.StatusOK {
		tag = "error"
	}

	c.JSON(res.Code, gin.H{
		tag: res.Body,
	})
}
