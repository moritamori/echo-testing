package repo

import (
	"github.com/moritamori/echo-testing/model"
	"gorm.io/gorm"
)

type BookRepoImpl struct {
	DB *gorm.DB
}

type BookRepo interface {
	FindByID(uint64) model.Book
	FindAll() []model.Book
}

func (br *BookRepoImpl) FindByID(id uint64) model.Book {
	book := model.Book{}
	br.DB.First(&book, id)
	return book
}

func (br *BookRepoImpl) FindAll() []model.Book {
	books := []model.Book{}
	br.DB.Find(&books)
	return books
}
