package services

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

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

func FetchAllNotes(notebookID string, pageStr string, user models.User) types.Response {
	const NOTES_PER_PAGE = 3
	var notebook models.Notebook
	initializers.DB.Find(&notebook, "id = ? AND user_id = ?", notebookID, user.ID)

	if notebook.ID == uuid.Nil {
		return types.Response{
			Code: http.StatusBadRequest,
			Body: "Cannot find the notebook.",
		}
	}

	page, err := strconv.Atoi(pageStr)

	if err != nil {
		return types.Response{
			Code: http.StatusBadRequest,
			Body: "Unable to fetch page.",
		}
	}

	var notesCount int64
	if err := initializers.DB.Table("notes").Count(&notesCount).Error; err != nil {
		return types.Response{
			Code: http.StatusInternalServerError,
			Body: "Failed to count notes",
		}
	}

	pageCount := int(math.Ceil(float64(notesCount) / float64(NOTES_PER_PAGE)))

	if pageCount == 0 {
		pageCount = 1
	}

	if page < 1 || page > pageCount {
		return types.Response{
			Code: http.StatusBadRequest,
			Body: "Invalid page.",
		}
	}

	offset := (page - 1) * NOTES_PER_PAGE

	var notes []models.Note
	result := initializers.DB.Limit(NOTES_PER_PAGE).Offset(offset).
		Find(&notes, "notebook_id = ?", notebookID)

	if result.Error != nil {
		return types.Response{
			Code: http.StatusInternalServerError,
			Body: "Failed to fetch notes.",
		}
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

	return types.Response{
		Code: http.StatusOK,
		Body: notes,
		Page: types.PaginatedData{
			PageCount: pageCount,
			CurrPage:  page,
			NextPage:  nextPage,
			PrevPage:  prevPage,
			Pages:     pages,
		},
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
