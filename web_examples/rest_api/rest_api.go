package main

/*
RESTful API Server for Vintage Jazz Records
--------------------------------------------

ENDPOINTS
----------
/albums
- GET – Get a list of all albums, returned as JSON.
- POST – Add a new album from request data sent as JSON.

/albums/:id
- GET – Get an album by its ID, returning the album data as JSON.

*/

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Album represents data about a record Album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// Responds with the list of all albums as JSON
// Creates JSOn from the slice of album structs, writing the JSON into the response
func getAlbums(c *gin.Context) {
	// Serializes the struct into JSON and adds it to the response
	c.IndentedJSON(http.StatusOK, albums)
}

// Adds an album from JSON received in the request body
func postAlbum(c *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	param_id := c.Param("id")

	for _, a := range albums {
		if a.ID == param_id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
	router.GET("/albums/:id", getAlbumByID)

	router.Run("localhost:8080")

	fmt.Println("-End of Main-")
}
