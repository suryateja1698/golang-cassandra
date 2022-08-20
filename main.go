package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/suryateja1698/golang-cassandra/db"
	"github.com/suryateja1698/golang-cassandra/handlers"
)

func main() {
	var port = "8080"

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	conn := db.Init()
	defer conn.Close()

	r := gin.New()
	handlers.Routes(r)

	// Start HTTP Server
	r.Run(":" + port)
}
