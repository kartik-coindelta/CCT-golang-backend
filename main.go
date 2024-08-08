package main

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/routes"
)

func main() {
	db.ConnectDB()

	r := routes.SetupRouter()
	r.Run(":8080")
}
