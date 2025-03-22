package models

import (
	"errors"
	"go_mysql_bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"name" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() (*Book, error) {
	if b.Name == "" || b.Author == "" {
		return nil, errors.New("name and author are required fields")
	}
	
	db.NewRecord(b)
	if err := db.Create(&b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

func GetAllBooks() ([]Book, error) {
	var Books []Book
	if err := db.Find(&Books).Error; err != nil {
		return nil, err
	}
	return Books, nil
}

func GetBookById(Id int64) (*Book, *gorm.DB, error) {
	var getBook Book
	result := db.Where("ID=?", Id).Find(&getBook)
	if result.Error != nil {
		return nil, nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil, errors.New("book not found")
	}
	return &getBook, db, nil
}

func DeleteBook(ID int64) (*Book, error) {
	var book Book
	// First check if the book exists
	if err := db.Where("ID=?", ID).First(&book).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}
	
	// Then delete it
	if err := db.Where("ID=?", ID).Delete(&Book{}).Error; err != nil {
		return nil, err
	}
	
	return &book, nil
}