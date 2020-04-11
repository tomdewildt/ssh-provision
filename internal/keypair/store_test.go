package keypair

import (
	"errors"
	"fmt"
	"os"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/tomdewildt/ssh-provision/pkg/ssh"
)

func TestNewStoreKeypairTask(t *testing.T) {
	client := ssh.ClientMock{
		NewSessionFunc:   func() (ssh.Session, error) { return nil, nil },
		ExecuteFunc:      func(cmd string) (string, error) { return "", nil },
		CloseSessionFunc: func() {},
	}
	task := NewStoreKeypairTask(&client).(storeKeypairTask)

	assert.Equal(t, &client, task.client, "Client should be equal to client")
	assert.Contains(t, task.inputSchema, "Server", "InputSchema should contain server")
	assert.Contains(t, task.inputSchema, "Username", "InputSchema should contain username")
	assert.Contains(t, task.inputSchema, "KeypairPath", "InputSchema should contain keypair path")
	assert.Contains(t, task.inputSchema, "PrivateKey", "InputSchema should contain privatekey")
	assert.Contains(t, task.inputSchema, "PublicKey", "InputSchema should contain publickey")
	assert.Contains(t, task.outputSchema, "Success", "OutputSchema should contain success")
}

func TestStoreKeypairExecuteEmptyInput(t *testing.T) {
	client := ssh.ClientMock{
		NewSessionFunc:   func() (ssh.Session, error) { return nil, nil },
		ExecuteFunc:      func(cmd string) (string, error) { return "", nil },
		CloseSessionFunc: func() {},
	}
	task := NewStoreKeypairTask(&client).(storeKeypairTask)

	out, err := task.Execute(nil)

	assert.Nil(t, out, "Out should be nil")
	assert.EqualError(t, err, "invalid input", "Error should be invalid input")
}

func TestStoreKeypairExecuteInvalidInput(t *testing.T) {
	client := ssh.ClientMock{
		NewSessionFunc:   func() (ssh.Session, error) { return nil, nil },
		ExecuteFunc:      func(cmd string) (string, error) { return "", nil },
		CloseSessionFunc: func() {},
	}
	task := NewStoreKeypairTask(&client).(storeKeypairTask)
	input := map[string]interface{}{
		"key": "value",
	}

	out, err := task.Execute(input)

	assert.Nil(t, out, "Out should be nil")
	assert.EqualError(t, err, "input does not match schema", "Error should be input does not match schema")
}

func TestStoreKeypairExecuteStoreKeypairError(t *testing.T) {
	input := map[string]interface{}{
		"Server":      fmt.Sprintf("test-%d", time.Now().Unix()),
		"Username":    "johndoe",
		"KeypairPath": "/tmp",
		"PrivateKey":  []byte{112, 114, 105, 118, 97, 116, 101, 45, 107, 101, 121},
		"PublicKey":   []byte{112, 117, 98, 108, 105, 99, 45, 107, 101, 121},
	}
	path := path.Join(input["KeypairPath"].(string), input["Server"].(string))

	client := ssh.ClientMock{
		NewSessionFunc: func() (ssh.Session, error) { return nil, nil },
		ExecuteFunc: func(cmd string) (string, error) {
			if cmd == fmt.Sprintf(storeKeypairCommand, input["PublicKey"], input["Username"]) {
				return "", errors.New("command error")
			}
			return "", nil
		},
		CloseSessionFunc: func() {},
	}
	task := NewStoreKeypairTask(&client).(storeKeypairTask)

	out, err := task.Execute(input)
	defer func() {
		os.Remove(path)
		os.Remove(path + ".pub")
	}()

	assert.Nil(t, out, "Out should be nil")
	assert.EqualError(t, err, "command error", "Error should be command error")
}

func TestStoreKeypairExecuteSetPermissionsError(t *testing.T) {
	input := map[string]interface{}{
		"Server":      fmt.Sprintf("test-%d", time.Now().Unix()),
		"Username":    "johndoe",
		"KeypairPath": "/tmp",
		"PrivateKey":  []byte{112, 114, 105, 118, 97, 116, 101, 45, 107, 101, 121},
		"PublicKey":   []byte{112, 117, 98, 108, 105, 99, 45, 107, 101, 121},
	}
	path := path.Join(input["KeypairPath"].(string), input["Server"].(string))

	client := ssh.ClientMock{
		NewSessionFunc: func() (ssh.Session, error) { return nil, nil },
		ExecuteFunc: func(cmd string) (string, error) {
			if cmd == fmt.Sprintf(setPermissionsCommand, input["Username"]) {
				return "", errors.New("command error")
			}
			return "", nil
		},
		CloseSessionFunc: func() {},
	}
	task := NewStoreKeypairTask(&client).(storeKeypairTask)

	out, err := task.Execute(input)
	defer func() {
		os.Remove(path)
		os.Remove(path + ".pub")
	}()

	assert.Nil(t, out, "Out should be nil")
	assert.EqualError(t, err, "command error", "Error should be command error")
}

func TestStoreKeypairExecute(t *testing.T) {
	input := map[string]interface{}{
		"Server":      fmt.Sprintf("test-%d", time.Now().Unix()),
		"Username":    "johndoe",
		"KeypairPath": "/tmp",
		"PrivateKey":  []byte{112, 114, 105, 118, 97, 116, 101, 45, 107, 101, 121},
		"PublicKey":   []byte{112, 117, 98, 108, 105, 99, 45, 107, 101, 121},
	}
	path := path.Join(input["KeypairPath"].(string), input["Server"].(string))

	client := ssh.ClientMock{
		NewSessionFunc:   func() (ssh.Session, error) { return nil, nil },
		ExecuteFunc:      func(cmd string) (string, error) { return "", nil },
		CloseSessionFunc: func() {},
	}
	task := NewStoreKeypairTask(&client).(storeKeypairTask)

	out, err := task.Execute(input)
	defer func() {
		os.Remove(path)
		os.Remove(path + ".pub")
	}()

	assert.Contains(t, out, "Success", "Out should contain success")
	assert.Equal(t, true, out["Success"], "Success should be true")
	assert.FileExists(t, path, "Private key should exist")
	assert.FileExists(t, path+".pub", "Public key should exist")
	assert.Nil(t, err, "Error should be nil")
}
