package command

type NewWindow struct {
	baseCommand
}

func NewWindowCommand(options ...string) *NewWindow {
	w := NewWindow{
		baseCommand: baseCommand{
			cmd: TMUX,
			args: []string{
				"new-window",
			},
		},
	}

	w.args = append(w.args, options...)

	return &w
}
