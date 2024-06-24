package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type Status string

const (
	IN_PROGRESS Status = "IN_PROGRESS"
	COMPLETED   Status = "COMPLETED"
)

type Todo struct {
	ID          uuid.UUID  `gorm:"column:id"`
	Title       string     `gorm:"column:title"`
	Description string     `gorm:"column:description"`
	Status      Status     `gorm:"column:status"`
	Image       []byte     `gorm:"column:image"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
}

func (Todo) TableName() string {
	return "tbl_todos"
}
