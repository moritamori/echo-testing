package repo

import (
	"regexp"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/moritamori/echo-testing/model"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// テストスイートの構造体
type BookRepoTestSuite struct {
	suite.Suite
	br   BookRepoImpl
	mock sqlmock.Sqlmock
}

// テストのセットアップ
func (suite *BookRepoTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mock = mock
	br := BookRepoImpl{}
	br.DB, _ = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	suite.br = br
}

// テスト終了時の処理（データベース接続のクローズ）
func (suite *BookRepoTestSuite) TearDownTest() {
	db, _ := suite.br.DB.DB()
	db.Close()
}

// テストスイートの実行
func TestBookRepoTestSuite(t *testing.T) {
	suite.Run(t, new(BookRepoTestSuite))
}

// FindByIDのテスト
func (suite *BookRepoTestSuite) TestFindByID() {
	suite.Run("find a book", func() {
		suite.mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "books" WHERE "books"."id" = $1 AND "books"."deleted_at" IS NULL ORDER BY "books"."id" LIMIT 1`)).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"title", "author"}).AddRow("Go言語の本", "誰か"))

		_, err := suite.br.FindByID(1)

		if err != nil {
			suite.Fail("Error発生", err)
		}
	})
}

// FindAllのテスト
func (suite *BookRepoTestSuite) TestFindAll() {
	suite.Run("find all book", func() {
		suite.mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "books" WHERE "books"."deleted_at" IS NULL`)).
			WillReturnRows(
				sqlmock.NewRows([]string{"title", "author"}).
					AddRow("Go言語の本", "誰か").
					AddRow("Go言語の本2", "誰か2"),
			)

		bks, err := suite.br.FindAll()

		if err != nil {
			suite.Fail("Error発生", err)
		}
		if len(bks) != 2 {
			suite.Fail("予期するレコードの件数と異なる")
		}
	})
}

// Createのテスト
func (suite *BookRepoTestSuite) TestCreate() {
	suite.Run("create a book", func() {
		newId := 1
		rows := sqlmock.NewRows([]string{"id"}).AddRow(newId)
		suite.mock.ExpectBegin()
		suite.mock.ExpectQuery(
			regexp.QuoteMeta(
				`INSERT INTO "books" ("created_at",` +
					`"updated_at","deleted_at","title",` +
					`"author") VALUES ($1,$2,$3,$4,$5) ` +
					`RETURNING "id"`),
		).WillReturnRows(rows)
		suite.mock.ExpectCommit()

		book := model.Book{
			Title:  "Go言語の本",
			Author: "誰か",
		}

		err := suite.br.Create(&book)

		if err != nil {
			suite.Fail("Error発生", err)
		}
		if book.ID != uint(newId) {
			suite.Fail("登録されるべきIDと異なっている")
		}
	})
}

// Saveのテスト
func (suite *BookRepoTestSuite) TestSave() {
	suite.Run("save a book", func() {
		suite.mock.ExpectBegin()
		suite.mock.ExpectExec(
			regexp.QuoteMeta(
				`UPDATE "books" SET "created_at"=$1,` +
					`"updated_at"=$2,"deleted_at"=$3,` +
					`"title"=$4,"author"=$5 ` +
					`WHERE "id" = $6`),
		).WillReturnResult(sqlmock.NewResult(1, 1))
		suite.mock.ExpectCommit()

		book := model.Book{
			Title:  "Go言語の本",
			Author: "誰か",
		}
		book.ID = 1

		err := suite.br.Save(&book)

		if err != nil {
			suite.Fail("Error発生", err)
		}
	})
}
