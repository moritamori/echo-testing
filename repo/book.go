package repo

import (
	"github.com/moritamori/echo-testing/model"
	"gorm.io/gorm"
)

type BookRepoImpl struct {
	DB *gorm.DB
}

type BookRepo interface {
	FindByID(uint64) (model.Book, error)
	FindAll() ([]model.Book, error)
	Create(book *model.Book) error
	Save(book *model.Book) error
}

func (br *BookRepoImpl) FindByID(id uint64) (model.Book, error) {
	b := model.Book{}
	err := br.DB.First(&b, id).Error
	return b, err
}

func (br *BookRepoImpl) FindAll() ([]model.Book, error) {
	bks := []model.Book{}
	err := br.DB.Find(&bks).Error
	return bks, err
}

func (br *BookRepoImpl) Create(book *model.Book) error {
	err := br.DB.Create(book).Error
	return err
}

func (br *BookRepoImpl) Save(book *model.Book) error {
	err := br.DB.Save(&book).Error
	return err
}
