package main

import (
	"os"
	"strconv"

	"github.com/coderockr/APIVagaFrontend/db"
	"github.com/coderockr/APIVagaFrontend/server"
)

// main ...
func main() {
	database := db.Connect()
	s := server.Setup(database)
	port := "8080"

	if p := os.Getenv("PORT"); p != "" {
		if _, err := strconv.Atoi(p); err == nil {
			port = p
		}
	}

	s.Run(":" + port)
}
