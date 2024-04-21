package model_test

import (
	"github.com/stretchr/testify/assert"
	"http-rest-api/internal/app/model"
	"testing"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser()

	assert.NoError(t, u.BeforeCreate())

	assert.NotEmpty(t, u.EncryptedPassword)
}

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser()
			},
			isValid: true,
		},
		{
			name: "invalid email",
			u: func() *model.User {
				u := model.TestUser()
				u.Email = "qwerty@mru"
				return u
			},
			isValid: false,
		},
		{
			name: "invalid password",
			u: func() *model.User {
				u := model.TestUser()
				u.Password = "qwerty"
				return u
			},
			isValid: false,
		},
		{
			name: "empty email",
			u: func() *model.User {
				u := model.TestUser()
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *model.User {
				u := model.TestUser()
				u.Password = ""
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}
