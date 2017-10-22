package command

type NewPane struct {
	TmuxCommand
}

func NewPaneCommand(options ...string) *NewPane {
	return &NewPane{
		TmuxCommand: NewTmuxCommand("split-window", options...),
	}
}
