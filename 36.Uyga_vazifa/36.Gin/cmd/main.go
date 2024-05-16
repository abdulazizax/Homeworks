package main

import (
	funk "h36/36.Gin/functions"

	"github.com/gin-gonic/gin"
)

func main() {
	db := funk.DBConnect()

	man := funk.Manager{
		DB: db,
	}

	man.CreateArtistTable()

	router := gin.Default()
	router.GET("/albums", funk.GetAlbums)
	router.GET("/albums/:id", funk.GetAlbumByID)
	router.POST("/albums", funk.PostAlbums)
	router.PUT("/albums/:id", funk.UpdateAlbums)

	router.Run("localhost:9000")
}
