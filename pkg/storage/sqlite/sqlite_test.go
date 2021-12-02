package sqlite

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	ts "github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

	s.repository = &repository{s.DB}
}

func (s *suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestStorageService(t *testing.T) {
	ts.Run(t, new(suite))
}
