package main

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "fmt"
)

// Define your model
type User struct {
    gorm.Model
    Name  string `gorm:"not null"`
    Email string `gorm:"unique;not null"`
    Age   int
}

func main() {
	// DATABASE_URL="postgresql://postgres:ahmad@localhost:5432/go_database?schema=public"
    // Connection string
    dsn := "host=localhost user=postgres dbname=go_database port=5432 password=ahmad sslmode=disable TimeZone=Asia/Shanghai"

    // Open connection
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }

    // Auto migrate schema
    db.AutoMigrate(&User{})

    // Create a user
    user := User{Name: "Alice", Email: "alice@example.com", Age: 30}
    result := db.Create(&user)
    if result.Error != nil {
        fmt.Println("Error creating user:", result.Error)
    }

    // Query user
    var foundUser User
    db.First(&foundUser, "email = ?", "alice@example.com")
    fmt.Printf("Found user: %+v\n", foundUser)

    // Update user
    db.Model(&foundUser).Update("Age", 31)

    // Delete user
    db.Delete(&foundUser)
}