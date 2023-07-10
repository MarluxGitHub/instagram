package main

import (
	"github.com/MarluxGitHub/instagram/cmd/instagram/routes"
	"github.com/MarluxGitHub/instagram/internal/database"
	"github.com/MarluxGitHub/instagram/internal/database/migration"
	"github.com/gin-gonic/gin"
)

func main() {
	migrateDatabase()

	r := gin.Default()

	routes.Setup(r)

	r.Run() // listen and serve on
}

func migrateDatabase() {
	db := database.Connect()
	migration.Migrate(db)
}