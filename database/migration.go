package database

import (
  "housy/models"
  "housy/pkg/mysql"
  "fmt"
)

// Automatic Migration if Running App
func RunMigration() {
  err := mysql.DB.AutoMigrate(&models.User{})

  if err != nil {
    fmt.Println(err)
    panic("Migration Failed")
  }

  fmt.Println("Migration Success")
}