package main

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestStringArray(t *testing.T) {
	sa := NewStringArray()
	assert.Equal(t, "", sa.String())
	err := sa.Set("foo")
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	assert.Equal(t, "foo", sa.String())
	err = sa.Set("bar")
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	assert.Equal(t, "foo,bar", sa.String())
}
