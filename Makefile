.PHONY: init clean run test build vm/start vm/stop vm/remove lint
.DEFAULT_GOAL := help

NAMESPACE := tomdewildt
NAME := ssh-provision

SERVER := vagrant
HOST := 10.20.0.5:22
ROOTUSERNAME := root
ROOTPASSWORD := vagrant
USERNAME := ${USER}
PASSWORD := password

help: ## Show this help
	@echo "${NAMESPACE}/${NAME}"
	@echo
	@fgrep -h "##" $(MAKEFILE_LIST) | \
	fgrep -v fgrep | sed -e 's/## */##/' | column -t -s##

##

init: ## Initialize the environment
	go mod download

clean: ## Clean the environment
	go mod tidy

##

run: ## Run the tool
	go run cmd/ssh-provision/ssh-provision.go \
		--server "${SERVER}"\
		--host "${HOST}"\
		--root-username "${ROOTUSERNAME}"\
		--root-password "${ROOTPASSWORD}"\
		--username "${USERNAME}"\
		--password "${PASSWORD}"

test: ## Run tests
	go test ./... 

##

build: ## Build the tool
	GOOS=linux go build \
		-o ssh-provision \
		-ldflags "-X main.Name=${NAME} -X main.Version=0.0.0" \
		cmd/ssh-provision/ssh-provision.go

##

vm/start: ## Start the virtual machine
	vagrant up

vm/stop: ## Stop the virtual machine
	vagrant halt

vm/remove: ## Remove the virtual machine
	vagrant destroy

##

lint: ## Run lint & syntax check
	go vet ./...
	gofmt -s -w .
