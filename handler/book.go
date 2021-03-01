package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/moritamori/echo-testing/model"
	"github.com/moritamori/echo-testing/repo"
)

type BookHandler struct {
	bookRepo repo.BookRepoImpl
}

type resultLists struct {
	Books []model.Book `json:"books"`
}

func NewBookHandler(br repo.BookRepoImpl) *BookHandler {
	return &BookHandler{br}
}

func (bh *BookHandler) GetIndex(c echo.Context) error {
	books := bh.bookRepo.FindAll()
	u := &resultLists{
		Books: books,
	}
	return c.JSON(http.StatusOK, u)
}

func (bh *BookHandler) GetDetail(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	b := bh.bookRepo.FindByID(id)
	return c.JSON(http.StatusOK, b)
}
