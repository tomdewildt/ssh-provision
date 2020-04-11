package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomdewildt/ssh-provision/pkg/ssh"
)

func TestNewRestartServiceTask(t *testing.T) {
	client := ssh.ClientMock{
		NewSessionFunc:   func() (ssh.Session, error) { return nil, nil },
		ExecuteFunc:      func(cmd string) (string, error) { return "", nil },
		CloseSessionFunc: func() {},
	}
	task := NewRestartServiceTask(&client).(restartServiceTask)

	assert.Equal(t, &client, task.client, "Client should be equal to client")
	assert.Equal(t, 0, len(task.inputSchema), "InputSchema should contain 0 elements")
	assert.Contains(t, task.outputSchema, "Success", "OutputSchema should contain success")
}

func TestRestartServiceExecuteEmptyInput(t *testing.T) {
	client := ssh.ClientMock{
		NewSessionFunc:   func() (ssh.Session, error) { return nil, nil },
		ExecuteFunc:      func(cmd string) (string, error) { return "", nil },
		CloseSessionFunc: func() {},
	}
	task := NewRestartServiceTask(&client).(restartServiceTask)

	out, err := task.Execute(nil)

	assert.Nil(t, out, "Out should be nil")
	assert.EqualError(t, err, "invalid input", "Error should be invalid input")
}

func TestRestartServiceExecuteRestartServiceError(t *testing.T) {
	client := ssh.ClientMock{
		NewSessionFunc:   func() (ssh.Session, error) { return nil, nil },
		ExecuteFunc:      func(cmd string) (string, error) { return "", errors.New("command error") },
		CloseSessionFunc: func() {},
	}
	task := NewRestartServiceTask(&client).(restartServiceTask)
	input := map[string]interface{}{}

	out, err := task.Execute(input)

	assert.Nil(t, out, "Out should be nil")
	assert.EqualError(t, err, "command error", "Error should be command error")
}

func TestRestartServiceExecute(t *testing.T) {
	client := ssh.ClientMock{
		NewSessionFunc:   func() (ssh.Session, error) { return nil, nil },
		ExecuteFunc:      func(cmd string) (string, error) { return "", nil },
		CloseSessionFunc: func() {},
	}
	task := NewRestartServiceTask(&client).(restartServiceTask)
	input := map[string]interface{}{}

	out, err := task.Execute(input)

	assert.Contains(t, out, "Success", "Out should contain success")
	assert.Equal(t, true, out["Success"], "Success should be true")
	assert.Nil(t, err, "Error should be nil")
}
