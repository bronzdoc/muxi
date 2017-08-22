package command

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
)

const TMUX = "tmux"

type TmuxCommand interface {
	Execute()
}

type baseCommand struct {
	cmd  string
	args []string
}

func (c *baseCommand) Execute() {
	if c.cmd == "" {
		fmt.Printf("Execute is not implemented for %v", reflect.TypeOf(c))
		return
	}

	cmd := exec.Command(c.cmd, c.args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
	}
}
