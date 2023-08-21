package services

import (
	"net/http"

	"github.com/Kunal-Patro/NoteTakingApp/dto"
	"github.com/Kunal-Patro/NoteTakingApp/initializers"
	"github.com/Kunal-Patro/NoteTakingApp/models"
	"github.com/Kunal-Patro/NoteTakingApp/types"
	"github.com/google/uuid"
)

func CreateNewNote(body *dto.NoteDTO, notebookID string, user models.User) types.Response {
	var notebook models.Notebook
	initializers.DB.Find(&notebook, "id = ? AND user_id = ?", notebookID, user.ID)

	if notebook.ID == uuid.Nil {
		return types.Response{
			Code: http.StatusBadRequest,
			Body: "Cannot find the notebook.",
		}
	}

	note := models.Note{
		Title:       body.Title,
		Description: body.Description,
		Notebook:    notebook,
	}

	result := initializers.DB.Create(&note)

	if result.Error != nil {
		return types.Response{
			Code: http.StatusInternalServerError,
			Body: "Failed to create note.",
		}
	}

	return types.Response{
		Code: http.StatusOK,
		Body: "New note created.",
	}
}

func FetchNote(noteID string, notebookID string, user models.User) types.Response {
	var notebook models.Notebook
	initializers.DB.Find(&notebook, "id = ? AND user_id = ?", notebookID, user.ID)

	if notebook.ID == uuid.Nil {
		return types.Response{
			Code: http.StatusBadRequest,
			Body: "Cannot find the notebook",
		}
	}

	var note models.Note
	initializers.DB.Find(&note, "id = ? AND notebook_id = ?", noteID, notebookID)

	if note.ID == uuid.Nil {
		return types.Response{
			Code: http.StatusBadRequest,
			Body: "Cannot find requested not in the notebook.",
		}
	}

	return types.Response{
		Code: http.StatusOK,
		Body: note,
	}
}

func AlterNote(body *dto.NoteDTO, noteID string, notebookID string, user models.User) types.Response {
	var notebook models.Notebook
	initializers.DB.Find(&notebook, "id = ? AND user_id = ?", notebookID, user.ID)

	if notebook.ID == uuid.Nil {
		return types.Response{
			Code: http.StatusBadRequest,
			Body: "Cannot find the notebook.",
		}
	}

	var note models.Note
	initializers.DB.Find(&note, "id = ? AND notebook_id = ?", noteID, notebookID)

	if note.ID == uuid.Nil {
		return types.Response{
			Code: http.StatusBadRequest,
			Body: "Cannot find requested note in the notebook",
		}
	}

	result := initializers.DB.Model(&note).Updates(models.Note{
		Title:       body.Title,
		Description: body.Description,
	})

	if result.Error != nil {
		return types.Response{
			Code: http.StatusInternalServerError,
			Body: "Failed to update note.",
		}
	}

	return types.Response{
		Code: http.StatusOK,
		Body: "Note updated.",
	}
}

func RemoveNote(noteID string, notebookID string, user models.User) types.Response {
	var notebook models.Notebook
	initializers.DB.Find(&notebook, "id = ? AND user_id = ?", notebookID, user.ID)

	if notebook.ID == uuid.Nil {
		return types.Response{
			Code: http.StatusBadRequest,
			Body: "Cannot find the notebook",
		}
	}

	var note models.Note
	initializers.DB.Find(&note, "id = ? AND notebook_id = ?", noteID, notebookID)

	if note.ID == uuid.Nil {
		return types.Response{
			Code: http.StatusBadRequest,
			Body: "Cannot find the requested note in notebook.",
		}
	}

	result := initializers.DB.Delete(&note)

	if result.Error != nil {
		return types.Response{
			Code: http.StatusInternalServerError,
			Body: "Failed to delete note",
		}
	}

	return types.Response{
		Code: http.StatusOK,
		Body: "Note deleted",
	}
}
