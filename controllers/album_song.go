package controller
//https://gorm.io/docs/query.html
//https://gorm.io/docs/create.html
import (
	"github.com/gin-gonic/gin"
	"github.com/haewonito/gin-gorm-rest/models"
	"github.com/haewonito/gin-gorm-rest/config"
)


func GetSongByAlbumId(c *gin.Context) {
	var album models.Album

	if err := config.DB.Preload("Songs").First(&album, c.Param("id")).Error; err != nil {
		panic(err)
	}

	var songs []models.Song
	for _, song := range album.Songs {
		songs = append(songs, song)
	}

	c.JSON(200, &songs)
}
