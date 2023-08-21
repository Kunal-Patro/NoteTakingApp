package services

import (
	"fmt"
	"net/http"

	"github.com/Kunal-Patro/NoteTakingApp/dto"
	"github.com/Kunal-Patro/NoteTakingApp/initializers"
	"github.com/Kunal-Patro/NoteTakingApp/models"
	"github.com/Kunal-Patro/NoteTakingApp/types"
	"github.com/google/uuid"
)

func CreateNotebook(body *dto.NotebookDTO, user models.User) types.Response {
	notebook := models.Notebook{
		Name: body.Name,
		User: user,
	}

	result := initializers.DB.Create(&notebook)

	if result.Error != nil {
		return types.Response{
			Code: http.StatusInternalServerError,
			Body: "Failed to create notebook",
		}
	}

	return types.Response{
		Code: http.StatusOK,
		Body: fmt.Sprintf(`Notebook by name "%v" created.`, notebook.Name),
	}
}

func FetchNotebook(notebookID string, user models.User) types.Response {
	var notebook models.Notebook
	initializers.DB.Find(&notebook, "id = ? AND user_id = ?", notebookID, user.ID)

	if notebook.ID == uuid.Nil {
		return types.Response{
			Code: http.StatusBadRequest,
			Body: "Cannot find the requested notebook.",
		}
	}

	return types.Response{
		Code: http.StatusOK,
		Body: notebook,
	}
}

func AlterNotebook(body *dto.NotebookDTO, notebookID string, user models.User) types.Response {
	var notebook models.Notebook
	initializers.DB.Find(&notebook, "id = ? AND user_id = ?", notebookID, user.ID)

	if notebook.ID == uuid.Nil {
		return types.Response{
			Code: http.StatusBadRequest,
			Body: "Cannot find requested notebook.",
		}
	}

	result := initializers.DB.Model(&notebook).Updates(models.Notebook{Name: body.Name})

	if result.Error != nil {
		return types.Response{
			Code: http.StatusInternalServerError,
			Body: "Failed to update notebook.",
		}
	}

	return types.Response{
		Code: http.StatusOK,
		Body: "Record Updated.",
	}
}

func RemoveNotebook(notebookID string, user models.User) types.Response {
	var notebook models.Notebook
	initializers.DB.Find(&notebook, "id = ? AND user_id = ?", notebookID, user.ID)

	if notebook.ID == uuid.Nil {
		return types.Response{
			Code: http.StatusBadRequest,
			Body: "Cannot find the requested notebook.",
		}
	}

	result := initializers.DB.Delete(&notebook)

	if result.Error != nil {
		return types.Response{
			Code: http.StatusInternalServerError,
			Body: "Failed to delete notebook",
		}
	}

	return types.Response{
		Code: http.StatusOK,
		Body: "Record deleted.",
	}
}
