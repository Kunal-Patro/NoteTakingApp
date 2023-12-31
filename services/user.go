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
	authKey := uuid.New()
	expiration := time.Now().Add(time.Hour * 24 * 30).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,    // subject
		"auth": authKey,    //authKey
		"exp":  expiration, // expiration
	})

	// Sign and get the complete encoded token as a sring using secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return types.Response{
			Code: http.StatusInternalServerError,
			Body: "Failed to generate token.",
		}, ""
	}

	newAuth := models.Auth{
		AuthID:          authKey,
		TokenExpiration: float64(expiration),
		User:            user,
	}
	result = initializers.DB.Create(&newAuth)

	if result.Error != nil {
		return types.Response{
			Code: http.StatusInternalServerError,
			Body: "Failed to create authentication",
		}, ""
	}

	return types.Response{
		Code: http.StatusOK,
		Body: "Logged In.",
	}, tokenString
}

func LogoutUser(auth models.Auth) types.Response {
	var authentication models.Auth
	result := initializers.DB.Find(&authentication, "auth_id = ?", auth.AuthID)

	if authentication.AuthID == uuid.Nil {
		return types.Response{
			Code: http.StatusBadRequest,
			Body: "Cannot find auth entry",
		}
	}

	if result.Error != nil {
		return types.Response{
			Code: http.StatusInternalServerError,
			Body: "Failed to fetch auth.",
		}
	}

	result = initializers.DB.Delete(&authentication)

	if result.Error != nil {
		return types.Response{
			Code: http.StatusInternalServerError,
			Body: "Failed to delete auth.",
		}
	}

	return types.Response{
		Code: http.StatusOK,
		Body: "Logged out.",
	}
}
