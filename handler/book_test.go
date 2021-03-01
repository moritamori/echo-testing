package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kinbiko/jsonassert"
	"github.com/labstack/echo"
	"github.com/moritamori/echo-testing/model"
	"github.com/stretchr/testify/assert"
)

type BookRepoStub struct{}

func (u *BookRepoStub) FindByID(id uint64) model.Book {
	return model.Book{
		Title:  "Go言語の本",
		Author: "誰か",
	}
}

func (u *BookRepoStub) FindAll() []model.Book {
	books := []model.Book{}
	books = append(books, model.Book{
		Title:  "Go言語の本",
		Author: "誰か",
	})
	books = append(books, model.Book{
		Title:  "Go言語の本2",
		Author: "誰か2",
	})
	return books
}

func TestGetDetail(t *testing.T) {
	ja := jsonassert.New(t)

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	brStub := &BookRepoStub{}
	h := NewBookHandler(brStub)

	if assert.NoError(t, h.GetDetail(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		ja.Assertf(
			rec.Body.String(), `
      {
        "ID": "<<PRESENCE>>",
        "CreatedAt": "<<PRESENCE>>",
        "UpdatedAt": "<<PRESENCE>>",
        "DeletedAt": null,
        "Title": "Go言語の本",
        "Author": "誰か"
			}`,
		)
	}
}

func TestGetIndex(t *testing.T) {
	ja := jsonassert.New(t)

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/books")

	brStub := &BookRepoStub{}
	h := NewBookHandler(brStub)

	if assert.NoError(t, h.GetIndex(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		ja.Assertf(
			rec.Body.String(), `
    		{
					"Books": [
						{
							"ID": "<<PRESENCE>>",
							"CreatedAt": "<<PRESENCE>>",
							"UpdatedAt": "<<PRESENCE>>",
							"DeletedAt": null,
							"Title": "Go言語の本",
							"Author": "誰か"
						},
						{
							"ID": "<<PRESENCE>>",
							"CreatedAt": "<<PRESENCE>>",
							"UpdatedAt": "<<PRESENCE>>",
							"DeletedAt": null,
							"Title": "Go言語の本2",
							"Author": "誰か2"
						}
					]
				}`,
		)
	}
}
