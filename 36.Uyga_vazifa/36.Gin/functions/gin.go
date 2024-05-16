package functions

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	db := DBConnect()

	man := Manager{
		DB: db,
	}

	albums := man.SelectAlbums()
	if len(albums) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Information is not available"})
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbums(c *gin.Context) {
	db := DBConnect()

	man := Manager{
		DB: db,
	}
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		log.Fatal(err)
	}

	result := man.InsertIntoAlbums(newAlbum)
	c.IndentedJSON(http.StatusCreated, result)
}

func GetAlbumByID(c *gin.Context) {
	db := DBConnect()
	man := Manager{
		DB: db,
	}
	id := c.Param("id")

	IdInt, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("Error convering string to int: ", err)
	}

	album, err := man.SelectAlbumsById(IdInt)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not founf"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func UpdateAlbums(c *gin.Context) {
	db := DBConnect()
	man := Manager{
		DB: db,
	}
	id := c.Param("id")

	IdInt, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("Error convering string to int: ", err)
	}

	var newAlbum Album
	if err := c.BindJSON(&newAlbum); err != nil {
		log.Fatal("Eror chiqdi: ", err)
	}

	result, err := man.AlterColumnById(newAlbum, IdInt)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not founf"})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}
