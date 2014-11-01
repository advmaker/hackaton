package main

import (
	"time"
	"github.com/eaigner/hood"
)

type Projects struct {
	id      	hood.Id
	email    	string
	password    string
	token    	string
	created_at 	time.Time
	deleted_at 	time.Time
	updated_at 	time.Time

}

func (m *M) CreateProjectsTable_1414838625_Up(hd *hood.Hood) {
	hd.CreateTable(&Projects{})
}

func (m *M) CreateProjectsTable_1414838625_Down(hd *hood.Hood) {
	hd.DropTableIfExists(&Projects{})
}
