package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type BaseCommand interface {
	Execute()
	PostHooks() []func()
	AddPostHook(func())
}

type TmuxCommand struct {
	cmd       string
	args      []string
	postHooks []func()
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

	for _, postHook := range c.postHooks {
		postHook()
	}
}

func (c *TmuxCommand) AddPostHook(hook func()) {
	c.postHooks = append(c.postHooks, hook)
}

func (c *TmuxCommand) PostHooks() []func() {
	return c.postHooks
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
