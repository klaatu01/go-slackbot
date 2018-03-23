package gobot

import (
	"strconv"
	"strings"
	"time"
)

type Uptime struct {
	name      string
	startTime time.Time
}

func (u Uptime) Name() string {
	return u.name
}

func (u Uptime) Init() {
	u.startTime = time.Now()
}

func (u Uptime) Run(args []string, message *RecvMessage) {
	text := ""
	if len(args) > 0 {
		switch strings.ToLower(args[0]) {
		case "mins", "minutes", "m":
			text = strconv.FormatFloat(time.Since(u.startTime).Minutes(), 'f', 2, 64)
		case "hours", "hrs", "h":
			text = strconv.FormatFloat(time.Since(u.startTime).Hours(), 'f', 2, 64)
		case "secs", "seconds", "s":
			text = strconv.FormatFloat(time.Since(u.startTime).Seconds(), 'f', 2, 64)
		case "help":
			text = "Command usage: uptime <mins|minutes|hours|hrs>"
		default:
			text = "Command usage: uptime <mins|minutes|hours|hrs>"
		}
	} else {
		text = time.Since(u.startTime).String()
	}

	response := &SendMessage{
		Text:    text,
		User:    message.User,
		Channel: message.Channel,
	}
	Response(response)
}

func NewUptime() Uptime {
	return Uptime{"uptime", time.Now()}
}
