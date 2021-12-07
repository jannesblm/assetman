package sqlite

import (
	"database/sql"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cmp307/assetman/pkg/auth"
	"github.com/cmp307/assetman/pkg/storage"
	"github.com/stretchr/testify/require"
	ts "github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"regexp"
	"testing"
	"time"
)

type suite struct {
	ts.Suite
	DB  *gorm.DB
	sql sqlmock.Sqlmock
	ar  *assetRepository
}

type mockAuthenticator struct {
	user *storage.User
}

func (m *mockAuthenticator) Authenticate(user string, password string) (storage.User, error) {
	m.user = &storage.User{
		Name:  user,
		Group: user,
	}

	return *m.user, nil
}

func (m *mockAuthenticator) Logout() {}

func (m *mockAuthenticator) GetUser() (storage.User, error) {
	if m.user == nil {
		return storage.User{}, auth.ErrNotAuthenticated
	}

	return *m.user, nil
}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
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

	db, s.sql, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.sql.ExpectQuery("select sqlite_version()").
		WillReturnRows(
			sqlmock.NewRows([]string{"sqlite_version()"}).
				AddRow(""))

	s.DB, err = gorm.Open(newSqliteDialector(db))
	require.NoError(s.T(), err)

	s.ar = &assetRepository{
		repository: repository{
			db: &DB{s.DB, &mockAuthenticator{}, true},
		},
	}

	s.ar.repository.db.InitialiseCallbacks()
}

func (s *suite) Test_QueryingOnUninitializedDbIsDisallowed() {
	s.ar.repository.db.Initialized = false
	_, err := s.ar.GetAllSoftware()
	s.ar.repository.db.Initialized = true

	require.ErrorIs(s.T(), err, ErrNotInitialized)
}

func (s *suite) Test_DisallowUnauthenticatedWrites() {
	err := s.ar.db.Create(storage.Asset{
		Name: "Test",
	}).Error

	require.ErrorIs(s.T(), err, ErrNoWritePermission)
}

func (s *suite) Test_DisallowReporterWrites() {
	s.ar.repository.db.auth.Authenticate("reporter", "reporter")

	err := s.ar.db.Create(storage.Asset{
		Name: "Test",
	}).Error

	require.ErrorIs(s.T(), err, ErrNoWritePermission)
}

func (s *suite) Test_AllowAdminWrites() {
	_, err := s.ar.repository.db.auth.Authenticate("admin", "admin")

	require.NoError(s.T(), err)

	s.sql.ExpectBegin()
	s.sql.ExpectExec(regexp.QuoteMeta("INSERT INTO `manufacturers`")).
		WithArgs(AnyTime{}, AnyTime{}, nil, "Test").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.sql.ExpectCommit()

	err = s.ar.db.Create(&storage.Manufacturer{
		Name: "Test",
	}).Error

	s.ar.repository.db.auth.Logout()
	require.NoError(s.T(), err)
}

func (s *suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.sql.ExpectationsWereMet())
}

func TestStorageService(t *testing.T) {
	ts.Run(t, new(suite))
}
