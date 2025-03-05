package main

import (
	"api-golang/database"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init(){
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

	databaseUrl := os.Getenv("DATABASE_URL");
	if databaseUrl == "" {
		log.Fatal("database connection url not found")
	}

	errDB := database.Connect(databaseUrl)
	if errDB != nil {
		log.Fatalf("Unable to connect to database: %v\n", errDB)
	} else {
		log.Println("DATABASE CONNECTED")
	}
}

func main(){
	r := gin.Default()
	var time time.Time

	r.GET("/", func(c *gin.Context) {
		time = database.GetTime(c)
		c.JSON(200, gin.H{
			"api": "golang",
			"now": time,
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	// listen and serve on 0.0.0.0:8080 (or "PORT" env var)
	r.Run()
}