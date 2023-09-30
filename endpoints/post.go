package endpoints

import (
	"net/http"

	"github.com/HPancieri/AlbumAPI/data"
	"github.com/HPancieri/AlbumAPI/models"
	"github.com/gin-gonic/gin"
)

func PostAlbum(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	data.Albums = append(data.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
