package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type album struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {
	initializeDB()
	//seedAlbums() -> Only enable this during development for first seed to populate db
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://albums-frontend.netlify.app/"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	//r.Use(cors.Default()) -> Allow from all origins

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getSpecificAlbum)
	router.POST("/albums", postAlbums)

	router.Run()
}
