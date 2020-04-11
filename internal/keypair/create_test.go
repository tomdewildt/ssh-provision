package keypair

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCreateKeypairTask(t *testing.T) {
	task := NewCreateKeypairTask().(createKeypairTask)

	assert.Contains(t, task.inputSchema, "KeypairBits", "InputSchema should contain keypair bits")
	assert.Contains(t, task.outputSchema, "Success", "OutputSchema should contain success")
	assert.Contains(t, task.outputSchema, "PrivateKey", "OutputSchema should contain privatekey")
	assert.Contains(t, task.outputSchema, "PublicKey", "OutputSchema should contain publickey")
}

func TestCreateKeypairExecuteEmptyInput(t *testing.T) {
	task := NewCreateKeypairTask().(createKeypairTask)

	out, err := task.Execute(nil)

	assert.Nil(t, out, "Out should be nil")
	assert.EqualError(t, err, "invalid input", "Error should be invalid input")
}

func TestCreateKeypairExecuteInvalidInput(t *testing.T) {
	task := NewCreateKeypairTask().(createKeypairTask)
	input := map[string]interface{}{
		"key": "value",
	}

	out, err := task.Execute(input)

	assert.Nil(t, out, "Out should be nil")
	assert.EqualError(t, err, "input does not match schema", "Error should be input does not match schema")
}

func TestCreateKeypairExecuteGenerateKeypairError(t *testing.T) {
	task := NewCreateKeypairTask().(createKeypairTask)
	input := map[string]interface{}{
		"KeypairBits": 3,
	}

	out, err := task.Execute(input)

	assert.Nil(t, out, "Out should be nil")
	assert.EqualError(t, err, "invalid bits count", "Error should be invalid bits count")
}

func TestCreateKeypairExecuteGenerateKeypair(t *testing.T) {
	task := NewCreateKeypairTask().(createKeypairTask)
	input := map[string]interface{}{
		"KeypairBits": 2048,
	}

	out, err := task.Execute(input)

	assert.Contains(t, out, "Success", "Out should contain success")
	assert.Contains(t, out, "PrivateKey", "Out should contain privatekey")
	assert.Contains(t, out, "PublicKey", "Out should contain publickey")
	assert.Equal(t, true, out["Success"], "Success should be true")
	assert.IsType(t, make([]byte, 0), out["PrivateKey"], "PrivateKey should be a byte slice")
	assert.IsType(t, make([]byte, 0), out["PublicKey"], "PublicKey should be a byte slice")
	assert.Nil(t, err, "Error should be nil")
}
