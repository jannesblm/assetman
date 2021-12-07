package fs

import (
	"context"
	"github.com/stretchr/testify/require"
	ts "github.com/stretchr/testify/suite"
	"testing"
)

type suite struct {
	ts.Suite
	fs *service
}

func (s *suite) SetupSuite() {
	ctx := context.WithValue(context.TODO(), "HomeDir", "C:\\Home")
	s.fs = NewService(ctx)
}

func (s *suite) Test_ReturnsCorrectDatabasePath() {
	require.EqualValues(s.T(), s.fs.GetDatabasePath(), "C:\\Home\\assets.db")
}

func (s *suite) Test_ReturnsCorrectHomePath() {
	require.EqualValues(s.T(), s.fs.GetAppHomePath(), "C:\\Home")
}

func (s *suite) Test_ReturnsCorrectBackupDirectory() {
	require.EqualValues(s.T(), s.fs.GetBackupDirectory(), "C:\\Home\\backups")
}

func TestFsService(t *testing.T) {
	ts.Run(t, new(suite))
}
