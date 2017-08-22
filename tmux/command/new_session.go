package command

type NewSession struct {
	baseCommand
}

func NewSessionCommand(sessionName string) *NewSession {
	return &NewSession{
		baseCommand: baseCommand{
			cmd: TMUX,
			args: []string{
				"rename-session",
				sessionName,
			},
		},
	}
}
