package models

import(
  "time"
)

type Achievement struct {
  Id int64
  ImageUrl string `sql:"type:varchar(255);"`
  Title string `sql:"type:varchar(128); not null;"`
  Description string
  ProgressLimit int16 `sql:"type:smallint;"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt time.Time

  ProjectId int64 `sql:"not null;"`
  Progresses []Progress
}
