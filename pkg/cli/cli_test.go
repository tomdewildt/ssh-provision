package cli

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCliNoRequiredFlags(t *testing.T) {
	var options *Options
	var arguments []string
	cli := NewCli("test", "0.0.0", func(opts Options, args []string) {
		options = &opts
		arguments = args
	})

	out := bytes.NewBufferString("")
	cli.SetOut(out)

	assert.Equal(t, "test", cli.Use, "Use should be equal to test")
	assert.Equal(t, "0.0.0", cli.Version, "Version should be equal to 0.0.0")

	err := cli.Execute()

	assert.Nil(t, options, "Options should be nil")
	assert.Nil(t, arguments, "Arguments should be nil")
	assert.Contains(t, err.Error(), "server", "Error should contain server")
	assert.Contains(t, err.Error(), "host", "Error should contain host")
	assert.Contains(t, err.Error(), "root-username", "Error should contain root-username")
	assert.Contains(t, err.Error(), "root-password", "Error should contain root-password")
	assert.Contains(t, err.Error(), "username", "Error should contain username")
	assert.Contains(t, err.Error(), "password", "Error should contain password")
}

func TestNewCliNoOptionalFlags(t *testing.T) {
	var options *Options
	var arguments []string
	cli := NewCli("test", "0.0.0", func(opts Options, args []string) {
		options = &opts
		arguments = args
	})

	out := bytes.NewBufferString("")
	cli.SetOut(out)
	cli.SetArgs([]string{
		"--server", "test",
		"--host", "127.0.0.1:22",
		"--root-username", "root",
		"--root-password", "password",
		"--username", "user",
		"--password", "password",
	})

	assert.Equal(t, "test", cli.Use, "Use should be equal to test")
	assert.Equal(t, "0.0.0", cli.Version, "Version should be equal to 0.0.0")

	err := cli.Execute()

	assert.Equal(t, false, options.DisablePasswordAccess, "DisablePasswordAccess should be false")
	assert.Equal(t, 2048, options.KeypairBits, "KeypairBits should be 2048")
	assert.Equal(t, fmt.Sprintf("%s/.ssh", os.Getenv("HOME")), options.KeypairPath, "KeypairPath should be $HOME/.ssh")
	assert.Equal(t, 0, len(arguments), "Arguments should be nil")
	assert.Nil(t, err, "Error should be nil")
}

func TestNewCli(t *testing.T) {
	var options *Options
	var arguments []string
	cli := NewCli("test", "0.0.0", func(opts Options, args []string) {
		options = &opts
		arguments = args
	})

	out := bytes.NewBufferString("")
	cli.SetOut(out)
	cli.SetArgs([]string{
		"--server", "test",
		"--host", "127.0.0.1:22",
		"--root-username", "root",
		"--root-password", "password",
		"--username", "user",
		"--password", "password",
		"--disable-password-access",
		"--keypair-bits", "1024",
		"--keypair-path", "/tmp/test",
	})

	assert.Equal(t, "test", cli.Use, "Use should be equal to test")
	assert.Equal(t, "0.0.0", cli.Version, "Version should be equal to 0.0.0")

	err := cli.Execute()

	assert.Equal(t, "test", options.Server, "Server should be test")
	assert.Equal(t, "127.0.0.1:22", options.Host, "Host should be 127.0.0.1:22")
	assert.Equal(t, "root", options.RootUsername, "RootUsername should be root")
	assert.Equal(t, "password", options.RootPassword, "RootPassword should be password")
	assert.Equal(t, "user", options.Username, "Username should be user")
	assert.Equal(t, "password", options.Password, "Password should be password")
	assert.Equal(t, true, options.DisablePasswordAccess, "DisablePasswordAccess should be true")
	assert.Equal(t, 1024, options.KeypairBits, "KeypairBits should be 1024")
	assert.Equal(t, "/tmp/test", options.KeypairPath, "KeypairPath should be /tmp/test")
	assert.Equal(t, 0, len(arguments), "Arguments should be nil")
	assert.Nil(t, err, "Error should be nil")
}
