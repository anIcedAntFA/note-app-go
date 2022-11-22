package main

import (
	"log"
	appctx "note_server/components/appctx"
	"note_server/middleware"
	notetransport "note_server/module/noteitem/transport"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("MYSQL_CONNECTION")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	db = db.Debug()

	appContext := appctx.NewAppContext(db)

	router := gin.Default()

	router.Use(middleware.Recover(appContext))

	v1 := router.Group("/v1")
	notes := v1.Group("/notes")

	notes.POST("", notetransport.HandleCreateNewNote(appContext))
	notes.GET("/:id", notetransport.HandleFindNote(appContext))
	notes.GET("", notetransport.HandleListNotes(appContext))
	notes.PATCH("/:id", notetransport.HandleUpdateNote(appContext))
	notes.DELETE("/:id", notetransport.HandleDeleteNote(appContext))

	router.Run()
}
