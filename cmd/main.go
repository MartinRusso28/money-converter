package main

import (
	"money-converter/internal/repository/sqlite"
	"money-converter/pkg/server"
	"os"
	"github.com/joho/godotenv"


	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	err := godotenv.Load(".env")
	
	if err != nil {
		panic("Error loading .env file")
	}

	db := sqlite.GetDatabase()

	server := moneysrv.GetMainEngine(db)

	err = server.Run(":" + os.Getenv("PORT"))

	if err != nil {
		panic("Error running the server")
	}
}
