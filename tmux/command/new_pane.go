package command

type NewPane struct {
	baseCommand
}

func NewPaneCommand() *NewPane {
	return &NewPane{
		baseCommand: baseCommand{
			cmd: TMUX,
			args: []string{
				"split-window",
			},
		},
	}
}
