package passwordaccess

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomdewildt/ssh-provision/pkg/ssh"
)

func TestNewDisablePasswordAccessTask(t *testing.T) {
	client := ssh.ClientMock{
		NewSessionFunc:   func() (ssh.Session, error) { return nil, nil },
		ExecuteFunc:      func(cmd string) (string, error) { return "", nil },
		CloseSessionFunc: func() {},
	}
	task := NewDisablePasswordAccessTask(&client).(disablePasswordAccessTask)

	assert.Equal(t, &client, task.client, "Client should be equal to client")
	assert.Contains(t, task.inputSchema, "DisablePasswordAccess", "InputSchema should contain disable password access")
	assert.Contains(t, task.outputSchema, "Success", "OutputSchema should contain success")
	assert.Contains(t, task.outputSchema, "Disabled", "OutputSchema should contain disabled")
}

func TestDisablePasswordAccessExecuteEmptyInput(t *testing.T) {
	client := ssh.ClientMock{
		NewSessionFunc:   func() (ssh.Session, error) { return nil, nil },
		ExecuteFunc:      func(cmd string) (string, error) { return "", nil },
		CloseSessionFunc: func() {},
	}
	task := NewDisablePasswordAccessTask(&client).(disablePasswordAccessTask)

	out, err := task.Execute(nil)

	assert.Nil(t, out, "Out should be nil")
	assert.EqualError(t, err, "invalid input", "Error should be invalid input")
}

func TestDisablePasswordAccessExecuteBackupConfigError(t *testing.T) {
	client := ssh.ClientMock{
		NewSessionFunc: func() (ssh.Session, error) { return nil, nil },
		ExecuteFunc: func(cmd string) (string, error) {
			if cmd == backupConfigCommand {
				return "", errors.New("command error")
			}
			return "", nil
		},
		CloseSessionFunc: func() {},
	}
	task := NewDisablePasswordAccessTask(&client).(disablePasswordAccessTask)
	input := map[string]interface{}{
		"DisablePasswordAccess": true,
	}

	out, err := task.Execute(input)

	assert.Nil(t, out, "Out should be nil")
	assert.EqualError(t, err, "command error", "Error should be command error")
}

func TestDisablePasswordAccessExecuteDisablePasswordAccessError(t *testing.T) {
	client := ssh.ClientMock{
		NewSessionFunc: func() (ssh.Session, error) { return nil, nil },
		ExecuteFunc: func(cmd string) (string, error) {
			if cmd == disablePasswordAccessCommand {
				return "", errors.New("command error")
			}
			return "", nil
		},
		CloseSessionFunc: func() {},
	}
	task := NewDisablePasswordAccessTask(&client).(disablePasswordAccessTask)
	input := map[string]interface{}{
		"DisablePasswordAccess": true,
	}

	out, err := task.Execute(input)

	assert.Nil(t, out, "Out should be nil")
	assert.EqualError(t, err, "command error", "Error should be command error")
}

func TestDisablePasswordAccessExecuteEnabled(t *testing.T) {
	client := ssh.ClientMock{
		NewSessionFunc:   func() (ssh.Session, error) { return nil, nil },
		ExecuteFunc:      func(cmd string) (string, error) { return "", nil },
		CloseSessionFunc: func() {},
	}
	task := NewDisablePasswordAccessTask(&client).(disablePasswordAccessTask)
	input := map[string]interface{}{
		"DisablePasswordAccess": true,
	}

	out, err := task.Execute(input)

	assert.Contains(t, out, "Success", "Out should contain success")
	assert.Contains(t, out, "Disabled", "Out should contain disabled")
	assert.Equal(t, true, out["Success"], "Success should be true")
	assert.Equal(t, true, out["Disabled"], "Disabled should be true")
	assert.Nil(t, err, "Error should be nil")
}

func TestDisablePasswordAccessExecuteDisabled(t *testing.T) {
	client := ssh.ClientMock{
		NewSessionFunc:   func() (ssh.Session, error) { return nil, nil },
		ExecuteFunc:      func(cmd string) (string, error) { return "", nil },
		CloseSessionFunc: func() {},
	}
	task := NewDisablePasswordAccessTask(&client).(disablePasswordAccessTask)
	input := map[string]interface{}{
		"DisablePasswordAccess": false,
	}

	out, err := task.Execute(input)

	assert.Contains(t, out, "Success", "Out should contain success")
	assert.Contains(t, out, "Disabled", "Out should contain disabled")
	assert.Equal(t, true, out["Success"], "Success should be true")
	assert.Equal(t, false, out["Disabled"], "Disabled should be false")
	assert.Nil(t, err, "Error should be nil")
}
