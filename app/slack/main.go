package main

import (
	"fmt"
	"github.com/nlopes/slack"
)

// set token starts with "xoxb-"
var api = slack.New("********", slack.OptionDebug(true))

func main(){
	a, b, err := api.PostMessage("*****",
		slack.MsgOptionBlocks(
			slack.NewSectionBlock(
				slack.NewTextBlockObject("plain_text", "hi :heart:", true, false),
				nil,
				nil),
			),
		)
	if err != nil {
		fmt.Printf("err: %+v\n", err)
		return
	}
	fmt.Printf("a: %+v, b: %+v\n", a, b)
}