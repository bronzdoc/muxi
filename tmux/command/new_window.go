package command

type NewWindow struct {
	baseCommand
}

func NewWindowCommand() *NewWindow {
	return &NewWindow{
		baseCommand: baseCommand{
			cmd: TMUX,
			args: []string{
				"new-window",
			},
		},
	}
}
