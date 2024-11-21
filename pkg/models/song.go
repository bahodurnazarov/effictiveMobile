package models

import "time"

type Song struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	Group     string     `json:"group_name" gorm:"column:group_name"`
	SongName  string     `json:"song_name" gorm:"column:song_name"`
	Link      string     `json:"link" gorm:"column:link"`
	Deleted   bool       `json:"deleted" gorm:"column:deleted"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"column:deleted_at"`
}
