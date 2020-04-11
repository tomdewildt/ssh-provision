package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapNoInput(t *testing.T) {
	out, err := Map(nil)

	assert.Nil(t, out, "Out should be nil")
	assert.EqualError(t, err, "invalid input", "Error should be invalid input")
}

func TestMapInvalidInput(t *testing.T) {
	out, err := Map("bogus")

	assert.Nil(t, out, "Out should be nil")
	assert.EqualError(t, err, "invalid input", "Error should be invalid input")
}

func TestMapEmptyInput(t *testing.T) {
	input := struct{}{}
	out, err := Map(input)

	assert.Equal(t, 0, len(out), "Out should be empty")
	assert.Nil(t, err, "Error should be nil")
}

func TestMapIntInput(t *testing.T) {
	input := struct {
		Key int
	}{Key: 1}
	out, err := Map(input)

	assert.Equal(t, 1, len(out), "Out should have one element")
	assert.Equal(t, 1, out["Key"], "Out element should be 1")
	assert.Nil(t, err, "Error should be nil")
}

func TestMapFloatInput(t *testing.T) {
	input := struct {
		Key float64
	}{Key: 1.5}
	out, err := Map(input)

	assert.Equal(t, 1, len(out), "Out should have one element")
	assert.Equal(t, 1.5, out["Key"], "Out element should be 1.5")
	assert.Nil(t, err, "Error should be nil")
}

func TestMapStringInput(t *testing.T) {
	input := struct {
		Key string
	}{Key: "test"}
	out, err := Map(input)

	assert.Equal(t, 1, len(out), "Out should have one element")
	assert.Equal(t, "test", out["Key"], "Out element should be test")
	assert.Nil(t, err, "Error should be nil")
}

func TestMapBoolInput(t *testing.T) {
	input := struct {
		Key bool
	}{Key: true}
	out, err := Map(input)

	assert.Equal(t, 1, len(out), "Out should contain 1 element")
	assert.Equal(t, true, out["Key"], "Out element should be true")
	assert.Nil(t, err, "Error should be nil")
}

func TestMap(t *testing.T) {
	input := struct {
		Int    int
		Float  float64
		String string
		Bool   bool
	}{
		Int:    1,
		Float:  1.5,
		String: "test",
		Bool:   true,
	}
	out, err := Map(input)

	assert.Equal(t, 4, len(out), "Out should contain 4 elements")
	assert.Equal(t, 1, out["Int"], "Out element should be 1")
	assert.Equal(t, 1.5, out["Float"], "Out element should be 1.5")
	assert.Equal(t, "test", out["String"], "Out element should be test")
	assert.Equal(t, true, out["Bool"], "Out element should be true")
	assert.Nil(t, err, "Error should be nil")
}
