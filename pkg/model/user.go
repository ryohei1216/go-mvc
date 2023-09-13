package model

import "time"

type User struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}
