package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	g.Run(os.Getenv("SERVER_PORT"))
}
