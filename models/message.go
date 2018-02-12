package models

import "github.com/jinzhu/gorm"

type Message struct {
	gorm.Model
	From    uint   `json:"from" binding:"required"`
	To      uint   `json:"to" binding:"required"`
	Message string `json:"message" `
}
