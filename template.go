package gobot

type Template struct {
	name string
}

func (t Template) Name() string {
	return t.name
}

func (t Template) Init() {

}

func (t Template) Run(args []string, message *RecvMessage) {
	text := ""

	response := &SendMessage{
		Text:    text,
		User:    message.User,
		Channel: message.Channel,
	}
	Response(response)
}

func NewTemplate() Template {
	return Status{"Status"}
}
