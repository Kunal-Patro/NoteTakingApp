package controllers

import (
	"fmt"
	"net/http"

	"github.com/Kunal-Patro/NoteTakingApp/dto"
	"github.com/Kunal-Patro/NoteTakingApp/initializers"
	"github.com/Kunal-Patro/NoteTakingApp/models"
	"github.com/Kunal-Patro/NoteTakingApp/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SignUp(c *gin.Context) {

	var body dto.UserDTO

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body",
		})
		return
	}

	res := services.CreateUser(&body)

	tag := "message"
	if res.Code != http.StatusOK {
		tag = "error"
	}

	c.JSON(res.Code, gin.H{
		tag: res.Body,
	})
}

func Login(c *gin.Context) {
	var body dto.LoginUserDTO

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	res, tokenString := services.LoginUser(&body)

	if tokenString != "" {
		// send that token back to client
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	}

	tag := "message"
	if res.Code != http.StatusOK {
		tag = "error"
	}
	c.JSON(res.Code, gin.H{
		tag: res.Body,
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%v is logged in successfully", user.(models.User).Email),
	})
}

func Logout(c *gin.Context) {
	auth, _ := c.Get("auth")

	var authentication models.Auth
	result := initializers.DB.Find(&authentication, "auth_id = ?", auth.(models.Auth).AuthID)

	if authentication.AuthID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot find auth entry.",
		})
		return
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch auth.",
		})
		return
	}

	result = initializers.DB.Delete(&authentication)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete auth.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out.",
	})
}
