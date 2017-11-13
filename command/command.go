package command

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/fatih/color"
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
	Stdin      io.Reader
	Stdout     bytes.Buffer
	Stderr     bytes.Buffer
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
	if err := c.runShell(c.cmd, c.cmdOptions); err != nil {
		fmt.Printf("Execute failded: %v\n", err)
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

func (c *TmuxCommand) runShell(command string, cmdOptions []string) error {
	internalCommand := "sh"
	internalCommandFlags := "-c"
	commandToExecute := fmt.Sprintf("%s %s", command, strings.Join(cmdOptions, " "))
	completeCommandString := fmt.Sprintf("%s %s %s", internalCommand, internalCommandFlags, commandToExecute)

	cmd := exec.Command(internalCommand, internalCommandFlags, commandToExecute)
	cmd.Stdout = &c.Stdout
	cmd.Stderr = &c.Stderr

	red := color.New(color.FgRed).SprintFunc()

	if err := cmd.Start(); err != nil {
		return errors.Wrapf(err, `could not start command %s`, completeCommandString)
	}

	if err := cmd.Wait(); err != nil {
		return errors.Wrapf(err, `command %s finished with errors: %s`, red(completeCommandString), red(c.Stderr.String()))
	}

	return nil
}
