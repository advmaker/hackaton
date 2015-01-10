package main

import (
	"github.com/advmaker/hood"
	"time"
)

type Projects struct {
	Id        hood.Id
	Secret    string `sql:"size(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (m *M) CreateProjectsTable_1420895518_Up(hd *hood.Hood) {
	hd.CreateTable(&Projects{})
}

func (m *M) CreateProjectsTable_1420895518_Down(hd *hood.Hood) {
	hd.DropTable(&Projects{})
}
