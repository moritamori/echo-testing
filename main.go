package main

import (
	"github.com/labstack/echo"
	"github.com/moritamori/echo-testing/db"
	"github.com/moritamori/echo-testing/handler"
	"github.com/moritamori/echo-testing/repo"
)

func main() {
	e := echo.New()

	d := db.DBConnect()
	br := &repo.BookRepoImpl{DB: d}
	h := handler.NewBookHandler(br)

	e.GET("/books", h.GetIndex)
	e.GET("/books/:id", h.GetDetail)

	e.Logger.Fatal(e.Start(":1324"))
}
