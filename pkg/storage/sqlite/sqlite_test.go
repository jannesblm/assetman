package sqlite

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cmp307/assetman/pkg/storage"
	"github.com/stretchr/testify/require"
	ts "github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

type suite struct {
	ts.Suite

	DB         *gorm.DB
	mock       sqlmock.Sqlmock
	repository *repository
}

func newSqliteDialector(conn gorm.ConnPool) gorm.Dialector {
	return &sqlite.Dialector{
		DriverName: sqlite.DriverName,
		DSN:        "",
		Conn:       conn,
	}
}

func (s *suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.mock.ExpectQuery("select sqlite_version()").
		WillReturnRows(
			sqlmock.NewRows([]string{"sqlite_version()"}).
				AddRow(""))

	s.DB, err = gorm.Open(newSqliteDialector(db))
	require.NoError(s.T(), err)

	s.repository = NewRepository(s.DB)
}

func (s *suite) Test_Asset_GetAll() {
	assetStub := storage.Asset{
		ID:   0,
		Name: "",
	}

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `assets` WHERE `assets`.`deleted_at` IS NULL")).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name"}).
				AddRow(assetStub.ID, assetStub.Name))

	_, err := s.repository.GetAll()
	require.NoError(s.T(), err)
}

func (s *suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestStorageService(t *testing.T) {
	ts.Run(t, new(suite))
}
