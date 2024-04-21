package sqlstore_test

import (
	"github.com/stretchr/testify/assert"
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/store"
	"http-rest-api/internal/app/store/sqlstore"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser()
	err := s.User().Create(u)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	someEmail := "qwerty@example.org"

	u0, err0 := s.User().FindByEmail(someEmail)
	assert.EqualError(t, err0, store.ErrRecordNotFound.Error())
	assert.Nil(t, u0)

	u := model.TestUser()
	u.Email = someEmail
	_ = s.User().Create(u)

	u1, err1 := s.User().FindByEmail(someEmail)
	assert.NoError(t, err1)
	assert.NotNil(t, u1)
}
