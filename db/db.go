package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string `json:"Name"`
	CPF    int    `json:"CPF"`
	Email  string `json:"Email"`
	Age    int    `json:"Age"`
	Active bool   `json:"Active"`
}

// github.com/mattn/go-sqlite3
func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Student{})

	return db
}

func AddStudent(student Student) {

	db := Init()

	if result := db.Create(&student); result.Error != nil {
		fmt.Println("Error to create student")
	}

	fmt.Println("Create student!")

}
