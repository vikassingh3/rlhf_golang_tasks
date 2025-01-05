package main  
import (  
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql" // Choose your database dialect
    "time"
)

type User struct {
    ID        uint       `gorm:"primary_key"`
    CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
    UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
    DeletedAt *time.Time `sql:"index"`
    Name      string     `gorm:"size:255"`
}

func main() {  
    // Initialize database connection as shown earlier
    db, err := gorm.Open("mysql", "user:root@/teastall?charset=utf8&parseTime=True&loc=Local&timeout=10s")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Automatic table migration
    db.AutoMigrate(&User{})

    // Create a new user
    user := User{Name: "John Doe"}
    db.Create(&user)

    // Find a user by ID
    var foundUser User
    db.First(&foundUser, user.ID)
    fmt.Println("Found user:", foundUser.Name)

    // Update a user's name
    db.Model(&foundUser).Update("Name", "Jane Doe")

    // Delete a user
    db.Delete(&foundUser)
}