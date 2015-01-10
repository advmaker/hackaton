package db

import (
	"github.com/advmaker/hood"
	"time"
)

type Achievements struct {
	Id            hood.Id
	ProjectId     int
	ImageUrl      string `sql:"size(255)"`
	Title         string `sql:"size(255)"`
	Description   string
	ProgressLimit int `sql:"size(2)"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type Users struct {
	Id        hood.Id
	Email     string `sql:"size(255)"`
	Password  string `sql:"size(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Projects struct {
	Id        hood.Id
	Secret    string `sql:"size(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Progresses struct {
	Id            hood.Id
	AchievementId int
	Progress      int `sql:"size(2)"`
	PlayerId      int
	PlayerExtra   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	UnlockedAt    time.Time
}
