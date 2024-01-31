package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID        string  `json: "id"`
	Title     string  `json: "title"`
	Artist    string  `json: "artist"`
	Price     float64 `json: "price"`
	Copmleted bool    `json: "completed"`
}

var albums = []Album{
	{ID: "1", Title: "Zetenegnaw shi", Artist: "Chubaw", Price: 19243.86, Copmleted: false},
	{ID: "2", Title: "Adey", Artist: "Bemnet", Price: 190243.86, Copmleted: true},
	{ID: "3", Title: "Besintu", Artist: "Alex", Price: 29243.86, Copmleted: false},
	{ID: "4", Title: "Zemen", Artist: "Solomon", Price: 49243.86, Copmleted: false},
}

func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(context *gin.Context) {
	var newAlbum Album

	if err := context.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumsByID(context *gin.Context) {
	id := context.Param("id")

	for _, a := range albums {
		if a.ID == id {
			context.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumsByID)

	router.Run("localhost:8000")
}
