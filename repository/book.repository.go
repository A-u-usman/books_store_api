package repository

import (
	entity "github.com/A-u-usman/books_strore_api/entities"
	"gorm.io/gorm"
)

// BookRepository is a  contract
type BookRepository interface {
	InsertBook(b entity.Book) entity.Book
	UpdateBook(b entity.Book) entity.Book
	DeleteBook(b entity.Book)
	AllBook() []entity.Book
	FindBookByID(bookID uint64) entity.Book
}

// // BookRepository is a  contract
// type categoryrepository interface {
// 	InsertCat(c models.Category) models.Category
// }

type bookConnection struct {
	connection *gorm.DB
}

// NewBOOkRepository create an instance BOOKRepository
func NewBOOkRepository(dbConn *gorm.DB) BookRepository {
	return &bookConnection{
		connection: dbConn,
	}
}

func (db *bookConnection) InsertBook(b entity.Book) entity.Book {
	db.connection.Save(&b)
	//db.connection.Preload("User").Find(&b)
	return b
}

//	func (db *bookConnection) InsertBook(b entity.Book) entity.Book {
//		db.connection.Save(&b)
//		return b
//	}
func (db *bookConnection) UpdateBook(b entity.Book) entity.Book {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *bookConnection) DeleteBook(b entity.Book) {
	db.connection.Delete(&b)
}

func (db *bookConnection) FindBookByID(bookID uint64) entity.Book {
	var book entity.Book
	db.connection.Preload("User").Find(&book, bookID)
	return book
}

func (db *bookConnection) AllBook() []entity.Book {
	var books []entity.Book
	db.connection.Preload("User").Find(&books)
	return books
}
