package endpoints

import (
	"net/http"

	"github.com/HPancieri/AlbumAPI/database"
	"github.com/gin-gonic/gin"
)

type result struct {
	ID     string  `json:"ID"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

func GetAlbums(c *gin.Context) {
	var res []result

	database.DB.Db.Raw(
		"SELECT id, title, artist, price FROM albums;",
	).Scan(&res)

	if len(res) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "there are no values in the database"})
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")
	var res result

	response := database.DB.Db.Raw(
		"SELECT id, title, artist, price FROM albums WHERE id = ?",
		id,
	).Scan(&res)

	if response.RowsAffected == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}
