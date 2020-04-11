package ssh

// ClientMock is an struct used to mock the Client object in the
// ssh-provision/pkg/ssh package. It has three parameters NewSessionFunc,
// ExecuteFunc and CloseSessionFunc. These functions will be called if
// the underlying NewSession(), Execute() and CloseSession() functions
// are called.
type ClientMock struct {
	NewSessionFunc   func() (Session, error)
	ExecuteFunc      func(cmd string) (string, error)
	CloseSessionFunc func()
}

// NewSession is used to open an session. The method takes no
// parameters and returns session and nil or nil and an error if one
// occurred.
func (m *ClientMock) NewSession() (Session, error) {
	return m.NewSessionFunc()
}

// Execute is used to run command on the ssh server. It takes a
// command as parameter and returns the output in string form and nil
// or an empty string and an error if one occurred.
func (m *ClientMock) Execute(cmd string) (string, error) {
	return m.ExecuteFunc(cmd)
}

// CloseSession is used to close an open session. The method takes no
// parameters.
func (m *ClientMock) CloseSession() {
	m.CloseSessionFunc()
}

// SessionMock is an struct used to mock the session object in the
// ssh-provision/pkg/ssh package. It has two parameters OutputFunc
// and CloseFunc of type func. These functions will be called if the
// underlying Output() and Close() functions are called.
type SessionMock struct {
	OutputFunc func(cmd string) ([]byte, error)
	CloseFunc  func() error
}

// Output is used to run a command on the remote server and collect
// it's output. It takes a command of type string as a parameter and
// returns an byte array and nil or nil and an error if one occurred.
func (m *SessionMock) Output(cmd string) ([]byte, error) {
	return m.OutputFunc(cmd)
}

// Close is used to close an open connection to the ssh server. It
// takes no parameters and returns nil or an error if one occurred.
func (m *SessionMock) Close() error {
	return m.CloseFunc()
}
