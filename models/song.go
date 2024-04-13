package models

import "gorm.io/gorm"
//
type Song struct {
	gorm.Model
	Id int 			`json:"ID" gorm:"primary_key"`
	Title string 	`json:"title"`
	AlbumId uint	`json:"album_id"`
}