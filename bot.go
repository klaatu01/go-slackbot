package gobot

import (
	"fmt"
	"strings"

	"github.com/nlopes/slack"
)

var botId string
var api *slack.Client
var rtm *slack.RTM

type RecvMessage struct {
	Channel *slack.Channel
	Text    string
	User    *slack.User
	Event   *slack.MessageEvent
}

type SendMessage struct {
	Channel *slack.Channel
	Text    string
	User    *slack.User
}

func StartBot(token string) {
	fmt.Println("Starting bot")
	InitCommands()
	api = slack.New(token)
	rtm = api.NewRTM()
	go rtm.ManageConnection()
Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				botId = ev.Info.User.ID
			case *slack.TeamJoinEvent:
				// Handle new user to client
			case *slack.MessageEvent:
				channel, err := api.GetChannelInfo(ev.Channel)
				if err != nil {
					//freakout
				}
				user, err := api.GetUserInfo(ev.User)
				if err != nil {
					//freakout
				}
				message := &RecvMessage{
					Channel: channel,
					Text:    ev.Text,
					User:    user,
					Event:   ev,
				}
				NewMessage(message)
			case *slack.ReactionAddedEvent:
				// Handle reaction added
			case *slack.ReactionRemovedEvent:
				// Handle reaction removed
			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())
			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop
			default:
				//fmt.Printf("Unknown error")
			}
		}
	}
}

func NewMessage(message *RecvMessage) {
	if strings.HasPrefix(message.Text, "<@"+botId+">") {
		message.Text = strings.TrimPrefix(message.Text, "<@"+botId+">")
		message.Text = strings.TrimSpace(message.Text)
		ParseMessage(message)
	}
}

func Response(response *SendMessage) {
	msg := rtm.NewOutgoingMessage(response.Text, response.Channel.ID)
	rtm.SendMessage(msg)
}

func ParseMessage(message *RecvMessage) {
	fmt.Println(message.Text)
	words := strings.Split(message.Text, " ")
	command := words[0]
	args := words[1:]
	RunCommand(command, args, message)
}
