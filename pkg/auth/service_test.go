package auth

import (
	"errors"
	"github.com/cmp307/assetman/pkg/storage"
	"github.com/stretchr/testify/require"
	ts "github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

type suite struct {
	ts.Suite
	ur   *mockUserRepository
	auth *service
}

type mockUserRepository struct {
	users map[string]storage.User
}

func (r *mockUserRepository) GetByName(name string) (storage.User, error) {
	if user, ok := r.users[name]; ok {
		return user, nil
	}

	return storage.User{}, errors.New("not found")
}

func (r *mockUserRepository) Save(user storage.User) error {
	r.users[user.Name] = user

	return nil
}

func (s *suite) SetupSuite() {
	s.ur = &mockUserRepository{
		users: make(map[string]storage.User, 1),
	}
	s.auth = NewService(s.ur)

	testPw, err := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)

	require.NoError(s.T(), err, "error hashing test password")

	s.ur.Save(storage.User{
		Name:     "user",
		Password: testPw,
		Group:    "",
	})
}

func (s *suite) Test_VerifiesPassword() {
	_, err := s.auth.Authenticate("user", "test")
	require.NoError(s.T(), err, "error authenticating user with correct password")

	s.auth.Logout()
}

func (s *suite) Test_FailsOnInvalidPassword() {
	_, err := s.auth.Authenticate("user", "wrong password")
	require.ErrorIs(s.T(), err, bcrypt.ErrMismatchedHashAndPassword)

	s.auth.Logout()
}

func (s *suite) Test_NotReturnsUserWhenUnauthenticated() {
	user, err := s.auth.GetUser()

	require.ErrorIs(s.T(), err, ErrNotAuthenticated)
	require.Empty(s.T(), user, "user struct is not empty")
}

func (s *suite) Test_ReturnsCorrectUserWhenAuthenticated() {
	_, err := s.auth.Authenticate("user", "test")

	require.NoError(s.T(), err, "error logging in with test user")

	user, err := s.auth.GetUser()

	require.NoError(s.T(), err, "error while getting user after authentication")
	require.EqualValues(s.T(), user.Name, "user")
}

func TestAuthService(t *testing.T) {
	ts.Run(t, new(suite))
}
