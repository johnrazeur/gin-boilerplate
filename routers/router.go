package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/johnrazeur/gin-boilerplate/controllers/user"
)

// Route make the routinf
func Route(app *gin.Engine) {
	app.POST("/signup", user.Signup)
	app.POST("/login", user.Login)
}
