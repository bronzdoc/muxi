package command

type NewPane struct {
	baseCommand
}

func NewPaneCommand(options ...string) *NewPane {
	p := NewPane{
		baseCommand: baseCommand{
			cmd: TMUX,
			args: []string{
				"split-window",
			},
		},
	}

	p.args = append(p.args, options...)

	return &p
}
