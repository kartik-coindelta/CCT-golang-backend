package main

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/routes"

	"github.com/joho/godotenv"
)

func main() {
	db.ConnectDB()
	err := godotenv.Load()
	if err != nil {
		// Handle the error
		panic(err)
	}
	r := routes.SetupRouter()
	r.Run(":8080")
}
