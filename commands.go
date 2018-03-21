package slackbot

var m map[string]func(string, []string, *RecvMessage) *SendMessage

func InitCommands() {
	m = make(map[string]func(string, []string, *RecvMessage) *SendMessage)
	AddCommand("henlo", Henlo)
}

func AddCommand(command string, function func(string, []string, *RecvMessage) *SendMessage) {
	m[command] = function
}

func RunCommand(command string, args []string, message *RecvMessage) *SendMessage {
	if m[command] != nil {
		return m[command](command, args, message)
	}
	return nil
}

func Henlo(command string, args []string, message *RecvMessage) *SendMessage {
	text := "Henlo there <@" + message.User.ID + ">!!!"
	response := &SendMessage{
		Channel: message.Channel,
		User:    message.User,
		Text:    text,
	}
	return response
}
