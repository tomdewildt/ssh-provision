package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Options is an struct used to represent all the options that the CLI
// tool has. It is passed to the run method when a new CLI is created.
type Options struct {
	Server                string
	Host                  string
	RootUsername          string
	RootPassword          string
	Username              string
	Password              string
	DisablePasswordAccess bool
	KeypairBits           int
	KeypairPath           string
}

// NewCli is used to create a new instance of a spf13/cobra Command.
// This method takes an name, version and run function as parameters.
// The run function will be called when arguments are parsed. The
// method returns an instance of a spf13/cobra Command.
func NewCli(name string, version string, run func(opts Options, args []string)) *cobra.Command {
	var opts Options
	cmd := &cobra.Command{
		Use:     name,
		Long:    "A simple tool to create ssh keys and distribute them to CentOS servers.",
		Version: version,
		Run: func(cmd *cobra.Command, args []string) {
			run(opts, args)
		},
	}

	cmd.Flags().StringVar(&opts.Server, "server", "", "name of the server")
	cmd.Flags().StringVar(&opts.Host, "host", "", "address of the server")
	cmd.Flags().StringVar(&opts.RootUsername, "root-username", "", "root username of the server")
	cmd.Flags().StringVar(&opts.RootPassword, "root-password", "", "root password of the server")
	cmd.Flags().StringVar(&opts.Username, "username", "", "username of user")
	cmd.Flags().StringVar(&opts.Password, "password", "", "password of user")
	cmd.Flags().BoolVar(&opts.DisablePasswordAccess, "disable-password-access", false, "disable ssh password access")
	cmd.Flags().IntVar(&opts.KeypairBits, "keypair-bits", 2048, "amount of bits in the keypair")
	cmd.Flags().StringVar(&opts.KeypairPath, "keypair-path", fmt.Sprintf("%s/.ssh", os.Getenv("HOME")), "location of the keypair")

	cmd.MarkFlagRequired("server")
	cmd.MarkFlagRequired("host")
	cmd.MarkFlagRequired("root-username")
	cmd.MarkFlagRequired("root-password")
	cmd.MarkFlagRequired("username")
	cmd.MarkFlagRequired("password")

	return cmd
}
