package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	gorm.Model
	Title string
	Price int
}

type Result struct {
	Total int
}

func dbInit() {
	db, err := gorm.Open("sqlite", "book.sqlite3")
	if err != nil {
		panic("You can't open DB (dbInsert())")
	}
	defer db.Close()
	db.AutoMigrate(&Book{})
}

// DB Create
func dbInsert(title string, price int) {
	db, err := gorm.Open("sqlite3", "book.sqlite3")
	if err != nil {
		panic("You can't open DB (dbInsert())")
	}
	defer db.Close()
	db.Create(&Book{Title: title, Price: price})
}

func dbUpdate(id int, title string, price int) {
	db, err := gorm.Open("sqlite3", "book.sqlite3")
	if err != nil {
		panic("You can't open DB (dbUpdate())")
	}
	defer db.Close()

	var book Book
	db.First(&book, id)
	book.Title = title
	book.Price = price
	db.Save(&book)
}

// DB Delete
func dbDelete(id int) {
	db, err := gorm.Open("sqlite3", "book.sqlite3")
	if err != nil {
		panic("You can't open DB (dbDelete())")
	}
	defer db.Close()
	var book Book
	db.First(&book, id)
	db.Delete(&book)
}

func dbGetAll() []Book {
	db, err := gorm.Open("sqlite3", "book.sqlite3")
	if err != nil {
		panic("You can't open DB (dbGetAll())")
	}
	defer db.Close()

	var books []Book
	db.Order("created_at desc").Find(&books)
	return books
}

func main() {

}
