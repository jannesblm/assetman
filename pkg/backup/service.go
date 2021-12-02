package backup

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

func (s *service) GetBackupDir() string {
	return filepath.FromSlash(s.GetAppHomePath() + "/backups")
}

func NewService(ctx context.Context) Service {
	s := &service{
		ctx: ctx,
	}

	err := os.MkdirAll(s.GetBackupDir(), os.ModePerm)

	if err != nil {
		panic("cannot initialise backup service")
	}

	return s
}

func (s *service) CreateBackup() (string, error) {
	db, err := os.OpenFile(s.GetDatabasePath(), os.O_RDONLY, os.ModePerm)
	defer db.Close()

	if err != nil {
		return "", err
	}

	name := filepath.FromSlash(s.GetBackupDir() + "/backup-" + strconv.FormatInt(time.Now().Unix(), 10) + ".db")
	backup, err := os.Create(name)

	defer backup.Close()

	if err != nil {
		return "", err
	}

	_, err = io.Copy(backup, db)

	return name, err
}

func (s *service) GetBackupList() ([]Backup, error) {
	var backups []Backup

	files, err := ioutil.ReadDir(s.GetBackupDir())

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
