package endpoints

import (
	"github.com/HPancieri/AlbumAPI/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Albums)
}

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, album := range data.Albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
