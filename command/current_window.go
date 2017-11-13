package command

import "os"

type CurrentWindow struct {
	TmuxCommand
}

func CurrentWindowCommand(options ...string) *CurrentWindow {
	cwd, _ := os.Getwd()
	return &CurrentWindow{
		TmuxCommand: NewTmuxCommand("rename-window", cwd),
	}
}
