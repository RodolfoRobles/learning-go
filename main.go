package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Album struct
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Album slice to seed record data.
var albums = []album{
	{ID: "1", Title: "Nevermind", Artist: "Nirvana", Price: 29.99},
	{ID: "2", Title: "NINJA", Artist: "NIN", Price: 49.99},
	{ID: "3", Title: "Suenos liquidos", Artist: "Mana", Price: 9.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.Run("localhost:8080")
}
