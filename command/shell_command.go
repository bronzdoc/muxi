package command

import "fmt"

type NewShell struct {
	TmuxCommand
}

func NewShellCommand(sessionName, cmd string) *NewShell {
	options := []string{
		"-t",
		sessionName,
		fmt.Sprintf("\"%s\"", cmd),
		"c-m",
	}

	return &NewShell{
		TmuxCommand: NewTmuxCommand("send-keys", options...),
	}
}
