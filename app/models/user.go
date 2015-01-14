package models

import(
  "time"
)

type User struct {
  Id int64
  Email string `sql:"type:varchar(50); not null;"`
  Password string `sql:"type:varchar(128); not null;"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt time.Time

  Projects []Project
}
