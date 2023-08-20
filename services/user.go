package services

import (
	"net/http"
	"os"
	"time"

	"github.com/Kunal-Patro/NoteTakingApp/dto"
	"github.com/Kunal-Patro/NoteTakingApp/initializers"
	"github.com/Kunal-Patro/NoteTakingApp/models"
	"github.com/Kunal-Patro/NoteTakingApp/types"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(body *dto.UserDTO) types.Response {

	// hash the passord
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		return types.Response{
			Code: http.StatusBadRequest,
			Body: "Failed to hash password",
		}
	}

	// Create User

	user := models.User{
		Email:        body.Email,
		PasswordHash: string(hash),
		FirstName:    body.FirstName,
		LastName:     body.LastName,
		Phone:        body.Phone,
		DateOfBirth:  body.DateOfBirth,
	}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		return types.Response{
			Code: http.StatusBadRequest,
			Body: "Failed to create user",
		}
	}

	// Respond
	return types.Response{
		Code: http.StatusOK,
		Body: "User Created!!.",
	}
}

func LoginUser(body *dto.LoginUserDTO) (types.Response, string) {

	// Get the user with email.
	var user models.User
	result := initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == uuid.Nil {
		return types.Response{
			Code: http.StatusBadRequest,
			Body: "Invalid email or password.",
		}, ""
	}

	if result.Error != nil {
		return types.Response{
			Code: http.StatusInternalServerError,
			Body: "Failed to fetch user data.",
		}, ""
	}

	// Compare provided password with saved user password
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password))

	if err != nil {
		return types.Response{
			Code: http.StatusBadRequest,
			Body: "Invalid email or password",
		}, ""
	}

	// Generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,                                    // subject
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(), // expiration
	})

	// Sign and get the complete encoded token as a sring using secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return types.Response{
			Code: http.StatusInternalServerError,
			Body: "Failed to generate token.",
		}, ""
	}

	return types.Response{
		Code: http.StatusOK,
		Body: "Logged In.",
	}, tokenString
}
