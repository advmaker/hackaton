package main

import (
	"github.com/eaigner/hood"
	"time"
)

type Achievements struct {
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

func (m *M) CreateAchievementsTable_1420891328_Up(hd *hood.Hood) {
	hd.CreateTable(&Achievements{})
}

func (m *M) CreateAchievementsTable_1420891328_Down(hd *hood.Hood) {
	hd.DropTable(&Achievements{})
}
