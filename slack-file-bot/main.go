package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func main() {
	err := os.Setenv("SLACK_BOT_TOKEN", "xoxb-6379935211040-6376879680193-V0PZHYsGm6X7OGUkREyXtIxi")
	if err != nil {
		return
	}
	err = os.Setenv("CHANNEL_ID", "C06A9MN97P1")
	if err != nil {
		return
	}
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"test1.pdf", "test1.txt", "test1.png"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}
