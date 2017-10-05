package command

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

const TMUX = "tmux"

type TmuxCommand interface {
	Execute() error
}

type baseCommand struct {
	cmd  string
	args []string
}

func (c *baseCommand) Execute() error {
	if c.cmd == "" {
		return fmt.Errorf("Execute is not implemented for %v", reflect.TypeOf(c))
	}
	return runShell(c.cmd, c.args)
}

func runShell(command string, args []string) error {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("%s %s", command, strings.Join(args, " ")))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("Error starting command %s", err)
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("Error waiting command: %s ", err)
	}

	return nil
}
