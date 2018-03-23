package gobot

var m map[string]func([]string, *RecvMessage)

type Command interface {
	Name() string
	Init()
	Run([]string, *RecvMessage)
}

func InitCommands() {
	m = make(map[string]func([]string, *RecvMessage))
	AddCommand(NewUptime())
}

func AddCommand(command Command) {
	command.Init()
	m[command.Name()] = command.Run
}

func RunCommand(command string, args []string, message *RecvMessage) {
	if m[command] != nil {
		m[command](args, message)
	}
}
