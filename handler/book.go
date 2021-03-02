package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/moritamori/echo-testing/model"
	"github.com/moritamori/echo-testing/repo"
)

type BookHandler struct {
	bookRepo repo.BookRepo
}

type resultLists struct {
	Books []model.Book `json:"Books"`
}

func NewBookHandler(br repo.BookRepo) *BookHandler {
	return &BookHandler{br}
}

func (bh *BookHandler) GetIndex(c echo.Context) error {
	bks, _ := bh.bookRepo.FindAll()
	b := &resultLists{
		Books: bks,
	}
	return c.JSON(http.StatusOK, b)
}

func (bh *BookHandler) GetDetail(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	b, _ := bh.bookRepo.FindByID(id)
	return c.JSON(http.StatusOK, b)
}

func (bh *BookHandler) Post(c echo.Context) error {
	t := c.FormValue("title")
	a := c.FormValue("author")
	b := model.Book{Title: t, Author: a}
	bh.bookRepo.Create(b)
	return c.JSON(http.StatusOK, b)
}

func (bh *BookHandler) Put(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	b, _ := bh.bookRepo.FindByID(id)

	t := c.FormValue("title")
	a := c.FormValue("author")

	b = model.Book{Title: t, Author: a}
	bh.bookRepo.Save(b)
	return c.JSON(http.StatusOK, b)
}
