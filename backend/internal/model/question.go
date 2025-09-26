package model

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	AuthorId uint   `json:"author_id"`
	Title    string `json:"title"`
	Content  string `json:"content" gorm:"type:varchar(2048)"`
	ImageUrl string `json:"image_url" gorm:"type:varchar(512)"`
	KeyPoint string `json:"key_point"`
}
