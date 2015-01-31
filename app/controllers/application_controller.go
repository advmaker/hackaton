package controllers

import(
  "os"
)

func ApplicationIndex() string {
  return os.Getenv("DB_ADAPTER")
}
