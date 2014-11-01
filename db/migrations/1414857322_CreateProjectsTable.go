package main

import (
	"github.com/eaigner/hood"
)

type Projects struct {
	Id      hood.Id
	First   string
	Last    string
}

func (m *M) CreateProjectsTable_1414857322_Up(hd *hood.Hood) {
	hd.CreateTable(&Projects{})
}

func (m *M) CreateProjectsTable_1414857322_Down(hd *hood.Hood) {
	hd.DropTable(&Projects{})
}
