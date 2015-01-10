package main

import (
	"github.com/eaigner/hood"
	"time"
)

type Users struct {
	Id        hood.Id
	Email     string `sql:"size(255)"`
	Password  string `sql:"size(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (m *M) CreateUsersTable_1420895106_Up(hd *hood.Hood) {
	hd.CreateTable(&Users{})
}

func (m *M) CreateUsersTable_1420895106_Down(hd *hood.Hood) {
	hd.DropTable(&Users{})
}
