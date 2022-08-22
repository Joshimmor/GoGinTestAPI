package models

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	input int `json:"input"`
}