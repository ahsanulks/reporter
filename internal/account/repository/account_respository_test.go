package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccountRepo(t *testing.T) {
	got := NewAccountRepo("test_input.csv", "test_output.csv")
	assert.Equal(t, &AccountRepo{
		InputFile:  "test_input.csv",
		OutputFile: "test_output.csv",
	}, got)
}
