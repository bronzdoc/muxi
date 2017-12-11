package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

type BaseCommand interface {
	Execute()
	PostHooks() []func()
	AddPostHook(func())
	Options() []string
}

type TmuxCommand struct {
	cmd        string
	cmdOptions []string
	options    []string
	postHooks  []func()
}

func NewTmuxCommand(tmuxCommand string, options ...string) TmuxCommand {
	t := TmuxCommand{
		cmd:        "tmux",
		cmdOptions: []string{tmuxCommand},
		options:    options,
	}

	t.cmdOptions = append(t.cmdOptions, options...)

	return t
}

func (c *TmuxCommand) Execute() {
	if err := runShell(c.cmd, c.cmdOptions); err != nil {
		panic(fmt.Sprintf("Execute failded %v", err))
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
func (c *TmuxCommand) Options() []string {
	return c.options
}

func runShell(command string, cmdOptions []string) error {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("%s %s", command, strings.Join(cmdOptions, " ")))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return errors.Wrap(err, "Error starting command")
	}

	if err := cmd.Wait(); err != nil {
		return errors.Wrap(err, "Error waiting command")
	}

	return nil
}
