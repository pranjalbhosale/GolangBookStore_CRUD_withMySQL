package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pranjalbhosale/Go_BookStore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model

	Name        string `gorm:"column:name" json:"name"`
	Author      string `gorm:"column:author" json:"author"`
	Publication string `gorm:"column:publication" json:"publication"`
}

func init() {

	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

//lets create all functions to talk to out DB

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	//The declaration var Books []Book creates a variable named Books that is a slice of Book structs.
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
	//The function returns a pointer to the getBook variable (which contains the retrieved book)

}
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID =?", ID).Delete(book)
	return book
}
