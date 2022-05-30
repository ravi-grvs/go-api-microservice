package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string `json:id`
	Title  string `json:title`
	Artist string `json:artist`
	Price  string `json:price`
}

var albums = []album{
	{ID: "1", Title: "album1", Artist: "ravi", Price: "100"},
	{ID: "2", Title: "album2", Artist: "kumar", Price: "100"},
	{ID: "3", Title: "album3", Artist: "kurva", Price: "100"},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context)  {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil{
		return
	}
	albums = append(albums, newAlbum);
	c.IndentedJSON(http.StatusOK,newAlbum)
}

func getAlbumById(c *gin.Context)  {
	id := c.Param("id")

	for _,a := range albums{
		if a.ID == id{
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound,gin.H{"message":"album not found"})
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.POST("/albums", postAlbum)
    router.GET("/albums/:id", getAlbumById)
    router.Run("localhost:8080")
}