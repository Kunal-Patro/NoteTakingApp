package controllers

import (
	"net/http"

	"github.com/Kunal-Patro/NoteTakingApp/dto"
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

	res := services.FetchAllNotebooks(pageStr, user.(models.User))

	tag := "notebooks"
	if res.Code != http.StatusOK {
		tag = "error"
	}

	c.JSON(res.Code, gin.H{
		tag:         res.Body,
		"page_data": res.Page, // page data will be {"page_count": 0,"curr_page": 0,"prev_page": "","next_page": "","pages": null} in case of error.
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
