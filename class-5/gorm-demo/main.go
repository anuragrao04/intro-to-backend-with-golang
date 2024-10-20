package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Classroom struct {
	gorm.Model
	Name     string
	Section  string
	Strength int
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("database connection screwed up")
	}

	// Migrate the schema
	db.AutoMigrate(&Classroom{})

	// create
	newClassroom := Classroom{
		Name:     "Introduction To Backend",
		Section:  "B",
		Strength: 45,
	}
	db.Create(&newClassroom)

	// read
	var retreivedClass Classroom
	db.First(&retreivedClass, 1)
	// find class with ID = 1

	// or

	db.First(&retreivedClass, "strength = ?", 45)
	// finds the first class with strength = 45

	fmt.Println(retreivedClass)

	// update
	db.Model(&retreivedClass).Update("Strength", 20)
	// update the retreived class to have strength 20

	db.Delete(&retreivedClass)
	// delete the retreived class

}
