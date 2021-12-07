package auth

import (
	"errors"
	"github.com/cmp307/assetman/pkg/storage"
	"golang.org/x/crypto/bcrypt"
)

var ErrNotAuthenticated = errors.New("not authenticated")

type Service interface {
	Authenticate(user string, password string) (storage.User, error)
	Logout()
	GetUser() (storage.User, error)
}

type service struct {
	r storage.UserRepository
	u *storage.User
}

func NewService(r storage.UserRepository) *service {
	return &service{
		r,
		nil,
	}
}

func (s *service) Authenticate(user string, password string) (storage.User, error) {
	u, err := s.r.GetByName(user)

	if err != nil {
		return storage.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword(u.Password, []byte(password)); err != nil {
		return storage.User{}, err
	}

	s.u = &u

	return u, nil
}

func (s *service) GetUser() (storage.User, error) {
	if s.u == nil {
		return storage.User{}, ErrNotAuthenticated
	}

	return *s.u, nil
}

func (s *service) Logout() {
	s.u = nil
}
