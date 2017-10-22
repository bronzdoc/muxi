package command

type SelectLayout struct {
	TmuxCommand
}

func NewSelectLayoutCommand(layout string) *SelectLayout {
	return &SelectLayout{
		TmuxCommand: NewTmuxCommand("select-layout", layout),
	}
}
