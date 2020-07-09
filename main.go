package main

import (
	"github.com/gin-gonic/gin"
	"github.com/johnrazeur/gin-boilerplate/config"
	"github.com/johnrazeur/gin-boilerplate/routers"
)

func main() {
	r := gin.Default()
	routers.Route(r)
	config.Init()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
