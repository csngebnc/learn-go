package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterStudentsUniqueByName(t *testing.T){
	input := []Student{
		{ Name: "Bob" },
		{ Name: "Alice" },
		{ Name: "Bob" },
	}

	expected := []Student{
		{ Name: "Bob" },
		{ Name: "Alice" },
	}

	assert.Equal(t, expected, FilterStudentsUniqueByName(input))
}