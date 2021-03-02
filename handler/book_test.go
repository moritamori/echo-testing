package handler

import (
	"net/http"
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/moritamori/echo-testing/model"
	"github.com/steinfletcher/apitest"
)

type BookRepoStub struct{}

func (u *BookRepoStub) FindByID(id uint64) (model.Book, error) {
	t, _ := time.Parse("2006-01-02", "2021-01-01")
	b := model.Book{Title: "Go言語の本", Author: "誰か"}
	b.ID = 1
	b.CreatedAt = t
	b.UpdatedAt = t
	return b, nil
}

func (u *BookRepoStub) FindAll() ([]model.Book, error) {
	bks := []model.Book{}
	t, _ := time.Parse("2006-01-02", "2021-01-01")

	bk1 := model.Book{Title: "Go言語の本", Author: "誰か"}
	bk1.ID = 1
	bk1.CreatedAt = t
	bk1.UpdatedAt = t
	bks = append(bks, bk1)

	b2 := model.Book{Title: "Go言語の本2", Author: "誰か2"}
	b2.ID = 2
	b2.CreatedAt = t
	b2.UpdatedAt = t
	bks = append(bks, b2)

	return bks, nil
}

func (u *BookRepoStub) Create(b model.Book) error {
	return nil
}

func (u *BookRepoStub) Save(b model.Book) error {
	return nil
}

func TestGetDetail(t *testing.T) {
	e := echo.New()
	brs := &BookRepoStub{}
	h := NewBookHandler(brs)
	e.GET("/books/:id", h.GetDetail)

	apitest.New().
		Handler(e).
		Get("/books/1").
		Expect(t).
		Body(`
			{
				"ID": 1,
				"CreatedAt": "2021-01-01T00:00:00Z",
				"UpdatedAt": "2021-01-01T00:00:00Z",
				"DeletedAt": null,
				"Title": "Go言語の本",
				"Author": "誰か"
			}
		`).
		Status(http.StatusOK).
		End()
}

func TestGetIndex(t *testing.T) {
	e := echo.New()
	brs := &BookRepoStub{}
	h := NewBookHandler(brs)
	e.GET("/books", h.GetIndex)

	apitest.New().
		Handler(e).
		Get("/books").
		Expect(t).
		Body(`
			{
				"Books": [
					{
						"ID": 1,
						"CreatedAt": "2021-01-01T00:00:00Z",
						"UpdatedAt": "2021-01-01T00:00:00Z",
						"DeletedAt": null,
						"Title": "Go言語の本",
						"Author": "誰か"
					},
					{
						"ID": 2,
						"CreatedAt": "2021-01-01T00:00:00Z",
						"UpdatedAt": "2021-01-01T00:00:00Z",
						"DeletedAt": null,
						"Title": "Go言語の本2",
						"Author": "誰か2"
					}
				]
			}
		`).
		Status(http.StatusOK).
		End()
}

func TestPost(t *testing.T) {
	e := echo.New()
	brs := &BookRepoStub{}
	h := NewBookHandler(brs)
	e.POST("/books", h.Post)

	apitest.New().
		Handler(e).
		Post("/books").
		FormData("title", "新規書籍名").
		FormData("author", "新規著者").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestPut(t *testing.T) {
	e := echo.New()
	brs := &BookRepoStub{}
	h := NewBookHandler(brs)
	e.PUT("/books/:id", h.Put)

	apitest.New().
		Handler(e).
		Put("/books/1").
		Expect(t).
		Status(http.StatusOK).
		End()
}
