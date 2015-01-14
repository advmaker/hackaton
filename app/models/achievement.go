package models

import(
  "time"
)

type Achievement struct {
  Id int64
  ProjectId int64
  ImageUrl string `sql:"type:varchar(255)"`
  Title string `sql:"type:varchar(128)"`
  Description string
  ProgressLimit int16 `sql:"type:smallint"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt time.Time
}
