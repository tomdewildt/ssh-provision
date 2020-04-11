package keypair

import (
	"fmt"
	"io/ioutil"
	"path"

	log "github.com/sirupsen/logrus"

	"github.com/tomdewildt/ssh-provision/pkg/ssh"
	"github.com/tomdewildt/ssh-provision/pkg/task"
	"github.com/tomdewildt/ssh-provision/pkg/validate"
)

const (
	storeKeypairCommand   = "echo '%s' > /home/%s/.ssh/authorized_keys"
	setPermissionsCommand = "chown %[1]s:%[1]s /home/%[1]s/.ssh/authorized_keys && chmod -R 644 /home/%[1]s/.ssh/authorized_keys"
)

type storeKeypairTask struct {
	client       ssh.Client
	inputSchema  map[string]string
	outputSchema map[string]string
}

// NewStoreKeypairTask is used to create a new storeKeypairTask.
// This method takes an ssh client as parameter and returns a task.
func NewStoreKeypairTask(client ssh.Client) task.Task {
	return storeKeypairTask{
		client: client,
		inputSchema: map[string]string{
			"Server":      "required,min=1",
			"Username":    "required,min=1",
			"KeypairPath": "required,min=1",
			"PrivateKey":  "required",
			"PublicKey":   "required",
		},
		outputSchema: map[string]string{
			"Success": "required",
		},
	}
}

func (t storeKeypairTask) Execute(input map[string]interface{}) (map[string]interface{}, error) {
	log.Info("Executing store keypair task")

	err := validate.Schema(t.inputSchema, input)
	if err != nil {
		return nil, err
	}

	output := map[string]interface{}{}
	path := path.Join(input["KeypairPath"].(string), input["Server"].(string))
	if err := ioutil.WriteFile(path, input["PrivateKey"].([]byte), 0600); err != nil {
		return nil, err
	}
	if err := ioutil.WriteFile(path+".pub", input["PublicKey"].([]byte), 0644); err != nil {
		return nil, err
	}

	_, err = t.client.Execute(fmt.Sprintf(storeKeypairCommand, input["PublicKey"], input["Username"]))
	if err != nil {
		return nil, err
	}
	_, err = t.client.Execute(fmt.Sprintf(setPermissionsCommand, input["Username"]))
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
