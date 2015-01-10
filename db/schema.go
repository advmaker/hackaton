package db

import (
	"github.com/eaigner/hood"
	"time"
)

type Achievement struct {
	Id            hood.Id
	ProjectId     int
	ImageUrl      string `sql:"size(255)"`
	Title         string `sql:"size(255)"`
	Description   string
	ProgressLimit int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}
