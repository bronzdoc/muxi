package command

type SelectLayout struct {
	baseCommand
}

func NewSelectLayoutCommand(layout string) *SelectLayout {
	return &SelectLayout{
		baseCommand: baseCommand{
			cmd: TMUX,
			args: []string{
				"select-layout",
				layout,
			},
		},
	}
}
