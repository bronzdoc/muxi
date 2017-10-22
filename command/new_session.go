package command

type newSessionCommand struct {
	TmuxCommand
}

func NewSessionCommand(sessionName string) *newSessionCommand {
	return &newSessionCommand{
		TmuxCommand: NewTmuxCommand("rename-session", sessionName),
	}
}
