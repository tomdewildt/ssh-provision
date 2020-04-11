package service

import (
	log "github.com/sirupsen/logrus"

	"github.com/tomdewildt/ssh-provision/pkg/ssh"
	"github.com/tomdewildt/ssh-provision/pkg/task"
	"github.com/tomdewildt/ssh-provision/pkg/validate"
)

const (
	restartServiceCommand = "service sshd restart"
)

type restartServiceTask struct {
	client       ssh.Client
	inputSchema  map[string]string
	outputSchema map[string]string
}

// NewRestartServiceTask is used to create a new restartServiceTask.
// This method takes an ssh client as parameter and returns a task.
func NewRestartServiceTask(client ssh.Client) task.Task {
	return restartServiceTask{
		client:      client,
		inputSchema: map[string]string{},
		outputSchema: map[string]string{
			"Success": "required",
		},
	}
}

func (t restartServiceTask) Execute(input map[string]interface{}) (map[string]interface{}, error) {
	log.Info("Executing restart service task")

	err := validate.Schema(t.inputSchema, input)
	if err != nil {
		return nil, err
	}

	output := map[string]interface{}{}
	_, err = t.client.Execute(restartServiceCommand)
	if err != nil {
		return nil, err
	}
	output["Success"] = true

	err = validate.Schema(t.outputSchema, output)
	if err != nil {
		return nil, err
	}

	return output, nil
}
