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

	runShell(c.cmd, c.args)
}

func runShell(command string, args []string) {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("%s %s", command, strings.Join(args, " ")))
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
