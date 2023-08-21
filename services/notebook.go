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

func FetchAllNotebooks(pageStr string, user models.User) types.Response {
	const NOTEBOOK_PER_PAGE = 3
	page, err := strconv.Atoi(pageStr)

	if err != nil {
		return types.Response{
			Code: http.StatusBadRequest,
			Body: "Unable to get page value.",
		}
	}

	var notebookCount int64
	if err := initializers.DB.Table("notebooks").Count(&notebookCount).Error; err != nil {
		return types.Response{
			Code: http.StatusInternalServerError,
			Body: "Failed to cout notebooks.",
		}
	}

	pageCount := int(math.Ceil(float64(notebookCount) / float64(NOTEBOOK_PER_PAGE)))

	if pageCount == 0 {
		pageCount = 1
	}
	if page < 1 || page > pageCount {
		return types.Response{
			Code: http.StatusBadRequest,
			Body: "Invalid Page",
		}
	}

	offset := (page - 1) * NOTEBOOK_PER_PAGE

	var notebooks []models.Notebook
	result := initializers.DB.Limit(NOTEBOOK_PER_PAGE).Offset(offset).
		Find(&notebooks, "user_id = ?", user.ID)

	if result.Error != nil {
		return types.Response{
			Code: http.StatusInternalServerError,
			Body: "Failed to fetch data.",
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

	for i := 0; i < pageCount; i++ {
		pages[i] = i + 1
	}

	return types.Response{
		Code: http.StatusOK,
		Body: notebooks,
		Page: types.PaginatedData{
			PageCount: pageCount,
			CurrPage:  page,
			NextPage:  nextPage,
			PrevPage:  prevPage,
			Pages:     pages,
		},
	}
}

func FetchNotebook(notebookID string, user models.User) types.Response {
	res, notebook := GetNotebookByUser(notebookID, user.ID)

	if res.Code != http.StatusOK {
		return res
	}

	return types.Response{
		Code: http.StatusOK,
		Body: notebook,
	}
}

func AlterNotebook(body *dto.NotebookDTO, notebookID string, user models.User) types.Response {
	res, notebook := GetNotebookByUser(notebookID, user.ID)

	if res.Code != http.StatusOK {
		return res
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
	res, notebook := GetNotebookByUser(notebookID, user.ID)

	if res.Code != http.StatusOK {
		return res
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

func GetNotebookByUser(notebookID string, userID uuid.UUID) (types.Response, models.Notebook) {
	var notebook models.Notebook
	initializers.DB.Find(&notebook, "id = ? AND user_id = ?", notebookID, userID)

	if notebook.ID == uuid.Nil {
		return types.Response{
			Code: http.StatusBadRequest,
			Body: "Cannot find the notebook",
		}, models.Notebook{}
	}

	return types.Response{
		Code: http.StatusOK,
		Body: "",
	}, notebook
}
