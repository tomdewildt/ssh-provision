package account

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomdewildt/ssh-provision/pkg/ssh"
)

func TestNewCreateAccountTask(t *testing.T) {
	client := ssh.ClientMock{
		NewSessionFunc:   func() (ssh.Session, error) { return nil, nil },
		ExecuteFunc:      func(cmd string) (string, error) { return "", nil },
		CloseSessionFunc: func() {},
	}
	task := NewCreateAccountTask(&client).(createAccountTask)

	assert.Equal(t, &client, task.client, "Client should be equal to client")
	assert.Contains(t, task.inputSchema, "Username", "InputSchema should contain username")
	assert.Contains(t, task.inputSchema, "Password", "InputSchema should contain password")
	assert.Contains(t, task.outputSchema, "Success", "OutputSchema should contain success")
	assert.Contains(t, task.outputSchema, "Created", "OutputSchema should contain created")
}

func TestCreateAccountExecuteEmptyInput(t *testing.T) {
	client := ssh.ClientMock{
		NewSessionFunc:   func() (ssh.Session, error) { return nil, nil },
		ExecuteFunc:      func(cmd string) (string, error) { return "", nil },
		CloseSessionFunc: func() {},
	}
	task := NewCreateAccountTask(&client).(createAccountTask)

	out, err := task.Execute(nil)

	assert.Nil(t, out, "Out should be nil")
	assert.EqualError(t, err, "invalid input", "Error should be invalid input")
}

func TestCreateAccountExecuteInvalidInput(t *testing.T) {
	client := ssh.ClientMock{
		NewSessionFunc:   func() (ssh.Session, error) { return nil, nil },
		ExecuteFunc:      func(cmd string) (string, error) { return "", nil },
		CloseSessionFunc: func() {},
	}
	task := NewCreateAccountTask(&client).(createAccountTask)
	input := map[string]interface{}{
		"key": "value",
	}

	out, err := task.Execute(input)

	assert.Nil(t, out, "Out should be nil")
	assert.EqualError(t, err, "input does not match schema", "Error should be input does not match schema")
}

func TestCreateAccountExecuteCreateUserError(t *testing.T) {
	client := ssh.ClientMock{
		NewSessionFunc: func() (ssh.Session, error) { return nil, nil },
		ExecuteFunc: func(cmd string) (string, error) {
			if cmd == fmt.Sprintf(checkUserCommand, "johndoe") {
				return "", errors.New("command error")
			}
			if cmd == fmt.Sprintf(createUserCommand, "johndoe", "password") {
				return "", errors.New("command error")
			}
			return "", nil
		},
		CloseSessionFunc: func() {},
	}
	task := NewCreateAccountTask(&client).(createAccountTask)
	input := map[string]interface{}{
		"Username": "johndoe",
		"Password": "password",
	}

	out, err := task.Execute(input)

	assert.Nil(t, out, "Out should be nil")
	assert.EqualError(t, err, "command error", "Error should be command error")
}

func TestCreateAccountExecuteCreateFolderError(t *testing.T) {
	client := ssh.ClientMock{
		NewSessionFunc: func() (ssh.Session, error) { return nil, nil },
		ExecuteFunc: func(cmd string) (string, error) {
			if cmd == fmt.Sprintf(checkUserCommand, "johndoe") {
				return "", errors.New("command error")
			}
			if cmd == fmt.Sprintf(createFolderCommand, "johndoe") {
				return "", errors.New("command error")
			}
			return "", nil
		},
		CloseSessionFunc: func() {},
	}
	task := NewCreateAccountTask(&client).(createAccountTask)
	input := map[string]interface{}{
		"Username": "johndoe",
		"Password": "password",
	}

	out, err := task.Execute(input)

	assert.Nil(t, out, "Out should be nil")
	assert.EqualError(t, err, "command error", "Error should be command error")
}

func TestCreateAccountExecuteSetPermissionsError(t *testing.T) {
	client := ssh.ClientMock{
		NewSessionFunc: func() (ssh.Session, error) { return nil, nil },
		ExecuteFunc: func(cmd string) (string, error) {
			if cmd == fmt.Sprintf(checkUserCommand, "johndoe") {
				return "", errors.New("command error")
			}
			if cmd == fmt.Sprintf(setPermissionsCommand, "johndoe") {
				return "", errors.New("command error")
			}
			return "", nil
		},
		CloseSessionFunc: func() {},
	}
	task := NewCreateAccountTask(&client).(createAccountTask)
	input := map[string]interface{}{
		"Username": "johndoe",
		"Password": "password",
	}

	out, err := task.Execute(input)

	assert.Nil(t, out, "Out should be nil")
	assert.EqualError(t, err, "command error", "Error should be command error")
}

func TestCreateAccountExecuteUserExist(t *testing.T) {
	client := ssh.ClientMock{
		NewSessionFunc: func() (ssh.Session, error) { return nil, nil },
		ExecuteFunc: func(cmd string) (string, error) {
			if cmd == fmt.Sprintf(checkUserCommand, "johndoe") {
				return "", errors.New("command error")
			}
			return "", nil
		},
		CloseSessionFunc: func() {},
	}
	task := NewCreateAccountTask(&client).(createAccountTask)
	input := map[string]interface{}{
		"Username": "johndoe",
		"Password": "password",
	}

	out, err := task.Execute(input)

	assert.Contains(t, out, "Success", "Out should contain success")
	assert.Contains(t, out, "Created", "Out should contain created")
	assert.Equal(t, true, out["Success"], "Success should be true")
	assert.Equal(t, true, out["Created"], "Created should be true")
	assert.Nil(t, err, "Error should be nil")
}

func TestCreateAccountExecuteUserNotExist(t *testing.T) {
	client := ssh.ClientMock{
		NewSessionFunc:   func() (ssh.Session, error) { return nil, nil },
		ExecuteFunc:      func(cmd string) (string, error) { return "", nil },
		CloseSessionFunc: func() {},
	}
	task := NewCreateAccountTask(&client).(createAccountTask)
	input := map[string]interface{}{
		"Username": "johndoe",
		"Password": "password",
	}

	out, err := task.Execute(input)

	assert.Contains(t, out, "Success", "Out should contain success")
	assert.Contains(t, out, "Created", "Out should contain created")
	assert.Equal(t, true, out["Success"], "Success should be true")
	assert.Equal(t, false, out["Created"], "Created should be false")
	assert.Nil(t, err, "Error should be nil")
}
