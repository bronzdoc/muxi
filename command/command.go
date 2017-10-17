package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type BaseCommand interface {
	Execute()
}

type TmuxCommand struct {
	cmd  string
	args []string
}

func NewTmuxCommand(tmuxCommand string, args ...string) TmuxCommand {
	t := TmuxCommand{
		cmd:  "tmux",
		args: []string{tmuxCommand},
	}

	t.args = append(t.args, args...)

	return t
}

func (c *TmuxCommand) Execute() {
	if err := runShell(c.cmd, c.args); err != nil {
		fmt.Printf("Execute failded %v", err)
	}
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
