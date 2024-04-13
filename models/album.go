package models

import "gorm.io/gorm"

//album has many songs
type Album struct {
	gorm.Model
    Id     int 		`json:"ID" gorm:"primary_key"`
    Title  string  	`json:"title"`
    Artist string	`json:"artist"`
    Price  float64 	`json:"price"`
	Songs []Song	`json:"songs" gorom:"foreignKey:AlbumRefer"`
}