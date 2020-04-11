package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchemaNoSchema(t *testing.T) {
	input := map[string]interface{}{
		"key": "value",
	}
	err := Schema(nil, input)

	assert.EqualError(t, err, "invalid schema", "Error should be invalid schema")
}

func TestSchemaNoInput(t *testing.T) {
	schema := map[string]string{
		"key": "required,alphanum,min=1",
	}
	err := Schema(schema, nil)

	assert.EqualError(t, err, "invalid input", "Error should be invalid input")
}

func TestSchemaValidInput(t *testing.T) {
	schema := map[string]string{
		"key": "required,alphanum,min=1",
	}
	input := map[string]interface{}{
		"key": "value",
	}
	err := Schema(schema, input)

	assert.Nil(t, err, "Error should be nil")
}

func TestSchemaInvalidInput(t *testing.T) {
	schema := map[string]string{
		"key": "required,alphanum,min=1",
	}
	input := map[string]interface{}{
		"key": 23,
	}
	err := Schema(schema, input)

	assert.EqualError(t, err, "input does not match schema", "Error should be input does not match schema")
}
