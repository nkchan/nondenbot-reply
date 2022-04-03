package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Response events.APIGatewayProxyResponse

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	line := Line{}
	err := line.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"))
	if err != nil {
		fmt.Println(err)
	}
	eve, err := ParseRequest(line.ChannelSecret, request)
	if err != nil {
		status := 200
		if err == linebot.ErrInvalidSignature {
			status = 400
		} else {
			status = 500
		}
		return events.APIGatewayProxyResponse{StatusCode: status}, errors.New("Bat Request")
	}
	line.EventRouter(eve)
	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
