package teststore

import (
	_ "github.com/lib/pq"
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/store"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{store: s, users: make(map[string]*model.User)}

	return s.userRepository
}
