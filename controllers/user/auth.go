package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johnrazeur/gin-boilerplate/services"
)

type signupForm struct {
	Email     string `form:"email" json:"email" binding:"required"`
	Username  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required,min=6,max=20"`
	Password2 string `form:"password2" json:"password2" binding:"required"`
}

type loginForm struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=20"`
}

// Signup a new user
func Signup(c *gin.Context) {
	var form signupForm

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if form.Password != form.Password2 {
		c.JSON(http.StatusOK, gin.H{"error": "Password does not match with conform password"})
		return
	}

	var userService services.UserService

	user, err := userService.Create(form.Email, form.Username, form.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// Login the user
func Login(c *gin.Context) {
	var form loginForm

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var userService services.UserService

	user, err := userService.Login(form.Email, form.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}
