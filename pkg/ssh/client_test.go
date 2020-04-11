package ssh

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/ssh"
)

func TestNewClientNoHost(t *testing.T) {
	client, err := NewClient("", "root", "password")

	assert.Nil(t, client, "Client should be nil")
	assert.EqualError(t, err, "invalid host", "Error should be invalid host")
}

func TestNewClientInvalidHost(t *testing.T) {
	client, err := NewClient("[{}]$%@&!:22", "root", "password")

	assert.Nil(t, client, "Client should be nil")
	assert.EqualError(t, err, "invalid host", "Error should be invalid host")
}

func TestNewClientNoPort(t *testing.T) {
	client, err := NewClient("127.0.0.1", "root", "password")

	assert.Nil(t, client, "Client should be nil")
	assert.EqualError(t, err, "invalid host", "Error should be invalid host")
}

func TestNewClientInvalidPort(t *testing.T) {
	client, err := NewClient("127.0.0.1:@$", "root", "password")

	assert.Nil(t, client, "Client should be nil")
	assert.EqualError(t, err, "invalid host", "Error should be invalid host")
}

func TestNewClientNoUsername(t *testing.T) {
	client, err := NewClient("127.0.0.1:22", "", "password")

	assert.Nil(t, client, "Client should be nil")
	assert.EqualError(t, err, "invalid username", "Error should be invalid username")
}

func TestNewClientInvalidUsername(t *testing.T) {
	client, err := NewClient("127.0.0.1:22", " ", "password")

	assert.Nil(t, client, "Client should be nil")
	assert.EqualError(t, err, "invalid username", "Error should be invalid username")
}

func TestNewClientNoPassword(t *testing.T) {
	client, err := NewClient("127.0.0.1:22", "root", "")

	assert.Nil(t, client, "Client should be nil")
	assert.EqualError(t, err, "invalid password", "Error should be invalid password")
}

func TestNewClientInvalidPassword(t *testing.T) {
	client, err := NewClient("127.0.0.1:22", "root", " ")

	assert.Nil(t, client, "Client should be nil")
	assert.EqualError(t, err, "invalid password", "Error should be invalid password")
}

func TestNewClientError(t *testing.T) {
	client, err := NewClient("127.0.0.1:22", "root", "password")

	assert.Nil(t, client, "Client should be nil")
	assert.NotNil(t, err, "Error should not be nil")
}

func TestNewSessionNoConnection(t *testing.T) {
	client := client{connection: nil, session: nil}

	out, err := client.NewSession()

	assert.Nil(t, client.session, "Session should be nil")
	assert.Nil(t, out, "Out should be nil")
	assert.EqualError(t, err, "connection closed", "Error should be connection closed")
}

func TestNewSessionExistingSession(t *testing.T) {
	connection := ssh.Client{}
	session := SessionMock{
		OutputFunc: func(cmd string) ([]byte, error) { return nil, nil },
		CloseFunc:  func() error { return nil },
	}
	client := client{connection: &connection, session: &session}

	out, err := client.NewSession()

	assert.Equal(t, &session, client.session, "Session should be equal to session")
	assert.Equal(t, &session, out, "Out should be equal to session")
	assert.Nil(t, err, "Error should be nil")
}

func TestExecuteNoCmd(t *testing.T) {
	connection := ssh.Client{}
	session := SessionMock{
		OutputFunc: func(cmd string) ([]byte, error) { return nil, nil },
		CloseFunc:  func() error { return nil },
	}
	client := client{connection: &connection, session: &session}

	out, err := client.Execute("")

	assert.Equal(t, "", out, "Output should be empty")
	assert.EqualError(t, err, "invalid cmd", "Error should be invalid cmd")
}

func TestExecuteInvalidCmd(t *testing.T) {
	connection := ssh.Client{}
	session := SessionMock{
		OutputFunc: func(cmd string) ([]byte, error) { return nil, nil },
		CloseFunc:  func() error { return nil },
	}
	client := client{connection: &connection, session: &session}

	out, err := client.Execute(" ")

	assert.Equal(t, "", out, "Output should be empty")
	assert.EqualError(t, err, "invalid cmd", "Error should be invalid cmd")
}

func TestExecuteNoConnection(t *testing.T) {
	client := client{connection: nil, session: nil}

	out, err := client.Execute("echo 'test'")

	assert.Equal(t, "", out, "Out should be empty")
	assert.EqualError(t, err, "session closed", "Error should be session closed")
}

func TestExecuteInvalidSessionData(t *testing.T) {
	connection := ssh.Client{}
	session := SessionMock{
		OutputFunc: func(cmd string) ([]byte, error) { return []byte{4, 8, 16, 32, 64, 128}, nil },
		CloseFunc:  func() error { return nil },
	}
	client := client{connection: &connection, session: &session}

	assert.Equal(t, &session, client.session, "Session not set")

	out, err := client.Execute("echo 'test'")

	assert.Nil(t, client.session, "Session should be nil")
	assert.Equal(t, "", out, "Output should be empty")
	assert.EqualError(t, err, "command error", "Error should be command error")
}

func TestExecuteSessionError(t *testing.T) {
	connection := ssh.Client{}
	session := SessionMock{
		OutputFunc: func(cmd string) ([]byte, error) { return nil, errors.New("command error") },
		CloseFunc:  func() error { return nil },
	}
	client := client{connection: &connection, session: &session}

	assert.Equal(t, &session, client.session, "Session not set")

	out, err := client.Execute("echo 'test'")

	assert.Nil(t, client.session, "Session should be nil")
	assert.Equal(t, "", out, "Output should be empty")
	assert.EqualError(t, err, "command error", "Error should be command error")
}

func TestExecute(t *testing.T) {
	connection := ssh.Client{}
	session := SessionMock{
		OutputFunc: func(cmd string) ([]byte, error) { return []byte{116, 101, 115, 116}, nil },
		CloseFunc:  func() error { return nil },
	}
	client := client{connection: &connection, session: &session}

	assert.Equal(t, &session, client.session, "Session not set")

	out, err := client.Execute("echo 'test'")

	assert.Nil(t, client.session, "Session should be nil")
	assert.Equal(t, "test", out, "Output should be test")
	assert.Nil(t, err, "Error should be nil")
}

func TestCloseSessionError(t *testing.T) {
	connection := ssh.Client{}
	session := SessionMock{
		OutputFunc: func(cmd string) ([]byte, error) { return nil, nil },
		CloseFunc:  func() error { return errors.New("session not closed") },
	}
	client := client{connection: &connection, session: &session}

	client.CloseSession()

	assert.Nil(t, client.session, "Session should be nil")
}

func TestCloseSession(t *testing.T) {
	connection := ssh.Client{}
	session := SessionMock{
		OutputFunc: func(cmd string) ([]byte, error) { return nil, nil },
		CloseFunc:  func() error { return nil },
	}
	client := client{connection: &connection, session: &session}

	client.CloseSession()

	assert.Nil(t, client.session, "Session should be nil")
}
