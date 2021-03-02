package repo

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// テストスイートの構造体
type BookRepositoryTestSuite struct {
	suite.Suite
	br   BookRepoImpl
	mock sqlmock.Sqlmock
}

// テストのセットアップ
func (suite *BookRepositoryTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mock = mock
	br := BookRepoImpl{}
	br.DB, _ = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	suite.br = br
}

// テスト終了時の処理（データベース接続のクローズ）
func (suite *BookRepositoryTestSuite) TearDownTest() {
	db, _ := suite.br.DB.DB()
	db.Close()
}

// テストスイートの実行
func TestBookRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(BookRepositoryTestSuite))
}

// Createのテスト
func (suite *BookRepositoryTestSuite) TestCreate() {
	suite.Run("get a book", func() {
	})
}
