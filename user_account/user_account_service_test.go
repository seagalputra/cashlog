package user_account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Add(first int, second int) int {
	return first + second
}

func TestAddMethod(t *testing.T) {
	result := Add(100, 200)
	assert.Equal(t, 300, result, "Should return sum of two number!")
}
