package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Response events.APIGatewayProxyResponse

func Handler(ctx context.Context) (string, error) {

	body, err := json.Marshal(map[string]interface{}{
		"message": "Go Serverless v1.0! Your function executed successfully! presents by nkchan",
	})
	if err != nil {
		return "error", err
	}

	log.Println(body)

	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
	var message_text string = "hoge"
	message := linebot.NewTextMessage(message_text)

	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
	return message_text, nil
}

func main() {
	lambda.Start(Handler)
}
