package endpoints

import (
	"github.com/HPancieri/AlbumAPI/data"
	"github.com/HPancieri/AlbumAPI/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostAlbum(c *gin.Context) {
	var newAlbum types.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	data.Albums = append(data.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
