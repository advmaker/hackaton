package models

import(
  "time"
)

type User struct {
  Id int64
  Email string `sql:"type:varchar(50);"`
  Password string `sql:"type:varchar(128);"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt time.Time
}
