package main

import (
	"github.com/advmaker/hood"
	"time"
)

type Progresses struct {
	Id             hood.Id
	Achievement_id int
	Progress       int
	Player_id      int
	Player_extra   string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	UnlockedAt     time.Time
}

func (m *M) CreateProgressesTable_1420895543_Up(hd *hood.Hood) {
	hd.CreateTable(&Progresses{})
}

func (m *M) CreateProgressesTable_1420895543_Down(hd *hood.Hood) {
	hd.DropTable(&Progresses{})
}
