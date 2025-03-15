package database

import "time"

type Message struct {
	ID        uint   `gorm:"PrimaryKey"`
	Username  string `gorm:"size:255"`
	Content   string
	CreatedAt time.Time
}
