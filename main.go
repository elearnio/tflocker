package main

import (
	"flag"
	"fmt"
	"os"

	tfBackendInit "github.com/hashicorp/terraform/backend/init"
	"github.com/hashicorp/terraform/command"
	"github.com/hashicorp/terraform/state"
	"github.com/mitchellh/cli"
)

func getState() (state.State, error) {
	// shitty workaround to prevent
	// cli to spam output
	os.Stderr.Close()

	// We don't want input either
	os.Stdin.Close()

	sm := command.StateMeta{
		Meta: command.Meta{
			Ui: &cli.BasicUi{
				Writer: os.Stderr,
				Reader: os.Stdin,
			},
		},
	}

	return sm.State()
}

func main() {
	var operation string
	flag.StringVar(&operation, "operation", "", "`operation` for lock info")
	flag.Parse()

	tfBackendInit.Init(nil)

	s, err := getState()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lockInfo := state.NewLockInfo()
	lockInfo.Operation = operation

	lockID, err := s.Lock(lockInfo)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(lockID)
}
