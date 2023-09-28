package main

import (
	"github.com/HPancieri/AlbumAPI/endpoints"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", endpoints.GetAlbums)
	router.GET("/albums/:id", endpoints.GetAlbumById)
	router.POST("/albums", endpoints.PostAlbum)

	router.Run("localhost:8080")
}
