package models

import "gorm.io/gorm"

type Zort struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	AuthorID    uint64 `json:"author_id"`
}
