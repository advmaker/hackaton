package models

import(
  "time"
)

type Progress struct {
  Id int64
  PlayerId int64 `sql:"not null;"`
  PlayerExtra string
  Progress int16 `sql:"type:smallint;"`
  UnlockedAt time.Time
  CreatedAt time.Time
  UpdatedAt time.Time

  AchievementId int64 `sql:"not null;"`
}
