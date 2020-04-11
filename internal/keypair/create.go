package keypair

import (
	log "github.com/sirupsen/logrus"

	"github.com/tomdewildt/ssh-provision/pkg/crypto"
	"github.com/tomdewildt/ssh-provision/pkg/task"
	"github.com/tomdewildt/ssh-provision/pkg/validate"
)

type createKeypairTask struct {
	inputSchema  map[string]string
	outputSchema map[string]string
}

// NewCreateKeypairTask is used to create a new createKeypairTask.
// This method takes no parameters and returns an task.
func NewCreateKeypairTask() task.Task {
	return createKeypairTask{
		inputSchema: map[string]string{
			"KeypairBits": "required,numeric,min=2",
		},
		outputSchema: map[string]string{
			"Success":    "required",
			"PrivateKey": "required,min=2",
			"PublicKey":  "required,min=2",
		},
	}
}

func (t createKeypairTask) Execute(input map[string]interface{}) (map[string]interface{}, error) {
	log.Info("Executing create keypair task")

	err := validate.Schema(t.inputSchema, input)
	if err != nil {
		return nil, err
	}

	output := map[string]interface{}{}
	privateKey, publicKey, err := crypto.GenerateKeypair(input["KeypairBits"].(int))
	if err != nil {
		return nil, err
	}
	output["Success"] = true
	output["PrivateKey"] = privateKey
	output["PublicKey"] = publicKey

	err = validate.Schema(t.outputSchema, output)
	if err != nil {
		return nil, err
	}

	return output, nil
}
