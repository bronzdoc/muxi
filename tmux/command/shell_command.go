package command

type NewShell struct {
	baseCommand
}

func NewShellCommand(sessionName, cmd string) *NewShell {
	return &NewShell{
		baseCommand: baseCommand{
			cmd: TMUX,
			args: []string{
				"send-keys",
				"-t",
				sessionName,
				cmd,
				"c-m",
			},
		},
	}
}
