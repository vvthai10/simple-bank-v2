package entity_test

import (
	"simple-bank-v2/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	u, err := entity.NewUser("vthai@gmail.com", "Van Thai", "123456")
	assert.Nil(t, err)
	assert.Equal(t, u.Name, "Van Thai")
	assert.Equal(t, u.Password, "123456")
}