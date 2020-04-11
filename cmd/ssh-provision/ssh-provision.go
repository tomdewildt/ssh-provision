package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/tomdewildt/ssh-provision/internal/account"
	"github.com/tomdewildt/ssh-provision/internal/keypair"
	"github.com/tomdewildt/ssh-provision/internal/passwordaccess"
	"github.com/tomdewildt/ssh-provision/internal/service"
	"github.com/tomdewildt/ssh-provision/pkg/cli"
	"github.com/tomdewildt/ssh-provision/pkg/ssh"
	"github.com/tomdewildt/ssh-provision/pkg/structs"
)

// Name of the binary
var Name string

// Version of the binary
var Version string

func main() {
	cli := cli.NewCli(Name, Version, run)
	if err := cli.Execute(); err != nil {
		os.Exit(0)
	}
}

func run(opts cli.Options, args []string) {
	// Create ssh client
	client, err := ssh.NewClient(opts.Host, opts.RootUsername, opts.RootPassword)
	if err != nil {
		log.Fatalln(err)
	}

	// Convert options to input
	input, err := structs.Map(opts)
	if err != nil {
		log.Fatalln(err)
	}

	// Run tasks
	createAccountTask := account.NewCreateAccountTask(client)
	_, err = createAccountTask.Execute(input)
	if err != nil {
		log.Fatalln(err)
	}

	createKeypairTask := keypair.NewCreateKeypairTask()
	out, err := createKeypairTask.Execute(input)
	if err != nil {
		log.Fatalln(err)
	}
	input["PrivateKey"] = out["PrivateKey"]
	input["PublicKey"] = out["PublicKey"]

	storeKeypairTask := keypair.NewStoreKeypairTask(client)
	_, err = storeKeypairTask.Execute(input)
	if err != nil {
		log.Fatalln(err)
	}

	disablePasswordAccessTask := passwordaccess.NewDisablePasswordAccessTask(client)
	_, err = disablePasswordAccessTask.Execute(input)
	if err != nil {
		log.Fatalln(err)
	}

	restartServiceTask := service.NewRestartServiceTask(client)
	_, err = restartServiceTask.Execute(input)
	if err != nil {
		log.Fatalln(err)
	}
}
