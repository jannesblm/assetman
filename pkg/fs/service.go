package fs

import (
	"context"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type Backup struct {
	Path     string
	Modified time.Time
	Size     int64
}

type Service interface {
	CreateBackup() (string, error)
	GetDatabasePath() string
}

type service struct {
	ctx context.Context
}

func (s *service) GetAppHomePath() string {
	return s.ctx.Value("HomeDir").(string)
}

func (s *service) GetDatabasePath() string {
	return filepath.FromSlash(s.GetAppHomePath() + "/assets.db")
}

func (s *service) GetBackupDirectory() string {
	return filepath.FromSlash(s.GetAppHomePath() + "/backups")
}

func NewService(ctx context.Context) *service {
	s := &service{
		ctx: ctx,
	}

	err := os.MkdirAll(s.GetBackupDirectory(), os.ModePerm)

	if err != nil {
		panic("cannot initialise fs service")
	}

	return s
}

// CreateBackup creates a backup at default backup path and returns its location.
func (s *service) CreateBackup() (string, error) {
	db, err := os.OpenFile(s.GetDatabasePath(), os.O_RDONLY, os.ModePerm)
	defer db.Close()

	if err != nil {
		return "", err
	}

	name := filepath.FromSlash(s.GetBackupDirectory() + "/backup-" + strconv.FormatInt(time.Now().Unix(), 10) + ".db")
	backup, err := os.Create(name)

	defer backup.Close()

	if err != nil {
		return "", err
	}

	_, err = io.Copy(backup, db)

	return name, err
}

// GetBackupList reads all files from the backup path.
func (s *service) GetBackupList() ([]Backup, error) {
	backups := make([]Backup, 0)

	files, err := ioutil.ReadDir(s.GetBackupDirectory())

	if err != nil {
		return []Backup{}, err
	}

	for _, file := range files {
		backups = append(backups, Backup{
			Path:     file.Name(),
			Modified: file.ModTime(),
			Size:     file.Size(),
		})
	}

	return backups, nil
}
