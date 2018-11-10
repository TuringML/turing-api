package main

import (
	"github.com/turing-ml/turing-api/db"
	"github.com/turing-ml/turing-api/server"
	"os"
	"strconv"
)

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
