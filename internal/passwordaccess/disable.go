package passwordaccess

import (
	log "github.com/sirupsen/logrus"

	"github.com/tomdewildt/ssh-provision/pkg/ssh"
	"github.com/tomdewildt/ssh-provision/pkg/task"
	"github.com/tomdewildt/ssh-provision/pkg/validate"
)

const (
	backupConfigCommand          = "cp /etc/ssh/sshd_config /etc/ssh/sshd_config.backup"
	disablePasswordAccessCommand = "sed -i 's|PasswordAuthentication yes|PasswordAuthentication no|g' /etc/ssh/sshd_config"
)

type disablePasswordAccessTask struct {
	client       ssh.Client
	inputSchema  map[string]string
	outputSchema map[string]string
}

// NewDisablePasswordAccessTask is used to create a new
// disablePasswordAccessTask. This method takes an ssh client as
// parameter and returns a task.
func NewDisablePasswordAccessTask(client ssh.Client) task.Task {
	return disablePasswordAccessTask{
		client: client,
		inputSchema: map[string]string{
			"DisablePasswordAccess": "",
		},
		outputSchema: map[string]string{
			"Success":  "required",
			"Disabled": "",
		},
	}
}

func (t disablePasswordAccessTask) Execute(input map[string]interface{}) (map[string]interface{}, error) {
	log.Info("Executing disable password access task")

	err := validate.Schema(t.inputSchema, input)
	if err != nil {
		return nil, err
	}

	output := map[string]interface{}{"Disabled": false}
	if input["DisablePasswordAccess"] != nil && input["DisablePasswordAccess"].(bool) {
		_, err = t.client.Execute(backupConfigCommand)
		if err != nil {
			return nil, err
		}

		_, err = t.client.Execute(disablePasswordAccessCommand)
		if err != nil {
			return nil, err
		}

		output["Disabled"] = true
	}
	output["Success"] = true

	err = validate.Schema(t.outputSchema, output)
	if err != nil {
		return nil, err
	}

	return output, nil
}
