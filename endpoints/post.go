package endpoints

import (
	"net/http"

	"github.com/HPancieri/AlbumAPI/database"
	"github.com/HPancieri/AlbumAPI/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostAlbum(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	newAlbum.ID = uuid.New()

	database.DB.Db.Exec(
		"INSERT INTO public.albums (id, created_at, updated_at, title, artist, price) VALUES (?, ?, ?, ?, ?, ?);",
		newAlbum.ID,
		newAlbum.CreatedAt,
		newAlbum.UpdatedAt,
		newAlbum.Title,
		newAlbum.Artist,
		newAlbum.Price,
	)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}
