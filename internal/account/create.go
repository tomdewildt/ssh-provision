package account

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/tomdewildt/ssh-provision/pkg/ssh"
	"github.com/tomdewildt/ssh-provision/pkg/task"
	"github.com/tomdewildt/ssh-provision/pkg/validate"
)

const (
	checkUserCommand      = "id -u %s"
	createUserCommand     = "adduser %[1]s && echo -e '%[2]s\n%[2]s' | passwd %[1]s && gpasswd -a %[1]s wheel"
	createFolderCommand   = "mkdir /home/%s/.ssh"
	setPermissionsCommand = "chown %[1]s:%[1]s /home/%[1]s/.ssh && chmod -R 700 /home/%[1]s/.ssh"
)

type createAccountTask struct {
	client       ssh.Client
	inputSchema  map[string]string
	outputSchema map[string]string
}

// NewCreateAccountTask is used to create a new createAccountTask.
// This method takes an ssh client as parameter and returns a task.
func NewCreateAccountTask(client ssh.Client) task.Task {
	return createAccountTask{
		client: client,
		inputSchema: map[string]string{
			"Username": "required,alphanum,min=1",
			"Password": "required,alphanum,min=1",
		},
		outputSchema: map[string]string{
			"Success": "required",
			"Created": "",
		},
	}
}

func (t createAccountTask) Execute(input map[string]interface{}) (map[string]interface{}, error) {
	log.Info("Executing create account task")

	err := validate.Schema(t.inputSchema, input)
	if err != nil {
		return nil, err
	}

	output := map[string]interface{}{"Created": false}
	_, err = t.client.Execute(fmt.Sprintf(checkUserCommand, input["Username"]))
	if err != nil {
		_, err = t.client.Execute(fmt.Sprintf(createUserCommand, input["Username"], input["Password"]))
		if err != nil {
			return nil, err
		}

		_, err = t.client.Execute(fmt.Sprintf(createFolderCommand, input["Username"]))
		if err != nil {
			return nil, err
		}

		_, err = t.client.Execute(fmt.Sprintf(setPermissionsCommand, input["Username"]))
		if err != nil {
			return nil, err
		}

		output["Created"] = true
	}
	output["Success"] = true

	err = validate.Schema(t.outputSchema, output)
	if err != nil {
		return nil, err
	}

	return output, nil
}
