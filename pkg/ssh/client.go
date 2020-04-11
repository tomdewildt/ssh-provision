package ssh

import (
	"errors"
	"regexp"
	"strings"
	"unicode/utf8"

	"golang.org/x/crypto/ssh"
)

type client struct {
	connection *ssh.Client
	session    Session
}

// Client is a simple interface wrapper around golang.org/x/crypto/ssh.
// It requires consumer to implement the NewSession() method, the
// Execute() method and the CloseSession() method.
type Client interface {
	// NewSession is used to open an session. The method takes no
	// parameters and returns session and nil or nil and an error if one
	// occurred.
	NewSession() (Session, error)

	// Execute is used to run command on the ssh server. It takes a
	// command as parameter and returns the output in string form and nil
	// or an empty string and an error if one occurred.
	Execute(cmd string) (string, error)

	// CloseSession is used to close an open session. The method takes no
	// parameters.
	CloseSession()
}

// Session is a simple interface wrapper around the session struct
// from golang.org/x/crypto/ssh. It requires consumer to implement
// the Output() method and the Close() method.
type Session interface {
	// Output is used to run a command on the remote server and collect
	// it's output. It takes a command of type string as a parameter and
	// returns an byte array and nil or nil and an error if one occurred.
	Output(cmd string) ([]byte, error)

	// Close is used to close an open connection to the ssh server. It
	// takes no parameters and returns nil or an error if one occurred.
	Close() error
}

// NewClient is used to create a instance of the client struct.
// It creates a ssh connection the to specified host. This method
// takes a host as parameter which should match the following
// pattern: 127.0.0.1:22 or example.com:22 and an username and
// password. It returns an instance of the client struct and nil
// or nil and an error if one occurred.
func NewClient(host string, username string, password string) (Client, error) {
	re, _ := regexp.Compile("(([a-z0-9]|[a-z0-9][a-z0-9\\-]*[a-z0-9])\\.)*([a-z0-9]|[a-z0-9][a-z0-9\\-]*[a-z0-9])(:[0-9]+)")
	if !re.MatchString(host) {
		return nil, errors.New("invalid host")
	}
	if strings.TrimSpace(username) == "" {
		return nil, errors.New("invalid username")
	}
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("invalid password")
	}

	config := &ssh.ClientConfig{
		User:            username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
	}
	connection, err := ssh.Dial("tcp", host, config)
	if err != nil {
		return nil, err
	}

	return &client{connection: connection, session: nil}, nil
}

func (c *client) NewSession() (Session, error) {
	if c.connection == nil {
		return nil, errors.New("connection closed")
	}

	if c.session == nil {
		session, err := c.connection.NewSession()
		if err != nil {
			return nil, err
		}
		c.session = session
		return session, nil
	}

	return c.session, nil
}

func (c *client) Execute(cmd string) (string, error) {
	if strings.TrimSpace(cmd) == "" {
		return "", errors.New("invalid cmd")
	}

	if _, err := c.NewSession(); err != nil {
		return "", errors.New("session closed")
	}
	defer c.CloseSession()

	raw, err := c.session.Output(cmd)
	if err != nil {
		return "", err
	}
	if !utf8.Valid(raw) {
		return "", errors.New("command error")
	}

	return string(raw), nil
}

func (c *client) CloseSession() {
	c.session.Close()
	c.session = nil
}
