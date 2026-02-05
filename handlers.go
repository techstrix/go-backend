package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getAlbums(c *gin.Context) {
	rows, err := db.Query("SELECT id, title, artist, price FROM albums")
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	albums := []album{}
	for rows.Next() {
		var a album
		rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
		albums = append(albums, a)
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func getSpecificAlbum(c *gin.Context) {
	id := c.Param("id")
	numId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid album id"})
		return
	}
	var alb album
	err = db.QueryRow("SELECT id, title, artist, price FROM albums WHERE id = ?", numId).Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, alb)

}
func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	_, err := db.Exec("INSERT INTO albums(title, artist, price) VALUES (?, ?, ?)",
		newAlbum.Title, newAlbum.Artist, newAlbum.Price)

	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(201, newAlbum)
}

func seedAlbums() {
	var albums = []album{
		{Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

	stmt, _ := db.Prepare("INSERT OR IGNORE INTO albums(title, artist, price) VALUES ( ?, ?, ?)")
	for _, album := range albums {
		stmt.Exec(album.Title, album.Artist, album.Price)

	}
}
