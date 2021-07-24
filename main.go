package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAllAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func addAlbum(c *gin.Context) {
	var newAlbum album

	err := c.BindJSON(&newAlbum)

	if err != nil {
		return
	}

	newAlbumId := len(albums) + 1

	newAlbum.ID = strconv.Itoa(newAlbumId)

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	var id = c.Param("id")

	var albumSearched album

	for _, a := range albums {
		if a.ID == id {
			albumSearched = a
			break
		}
	}

	if (albumSearched == album{}) {
		c.IndentedJSON(http.StatusNotFound, nil)
	} else {
		c.IndentedJSON(http.StatusOK, albumSearched)
	}
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAllAlbums)
	router.POST("/albums", addAlbum)

	router.GET("/albums/:id", getAlbumById)

	router.Run("localhost:3000")
}
