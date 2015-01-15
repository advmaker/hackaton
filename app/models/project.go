package models

import(
  "time"
)

type Project struct {
  Id int64
  Secret string `sql:"type:varchar(128);"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt time.Time

  UserId int64 `sql:"not null;"`
  Achievements []Achievement
}
