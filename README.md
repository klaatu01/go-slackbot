# GoBot : Basic bot for slack using Go

GoBot is a simple framework for creating a basic bot for slack.

#### Usage:

Entry point: _token is your slack-bot token._
```go
gobot.StartBot(token)
```

Example main function:
```go
import (
    "gobot"
    "os"
    )
func main(){
    gobot.StartBot(os.Args[1])
}
```

Creating a Command:
_..._

[nlopes Slack API](https://github.com/nlopes/slack)
