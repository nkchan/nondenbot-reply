package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/linebot"
)

type Response events.APIGatewayProxyResponse

func UnmarshalLineRequest(data []byte) (LineRequest, error) {
	var r LineRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *LineRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type LineRequest struct {
	Events      []Event `json:"events"`
	Destination string  `json:"destination"`
}

type Event struct {
	Type       string  `json:"type"`
	ReplyToken string  `json:"replyToken"`
	Source     Source  `json:"source"`
	Timestamp  int64   `json:"timestamp"`
	Message    Message `json:"message"`
}

type Message struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Text string `json:"text"`
}

type Source struct {
	UserID string `json:"userId"`
	GroupID string `json:"GroupId"`
	Type   string `json:"type"`
}

func (data Source) GetId() string{
	if data.Type == "user" {
		return data.UserID 
	} else if data.Type == "group" {
		return data.GroupID 
	}
	return "error"
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	bot, err := linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_ACCESS_TOKEN"),
	)

	if err != nil {
		log.Fatal(err)
	}
	log.Print(request.Headers)
	log.Print(request.Body)
	log.Print(bot.GetBotInfo().Do())
	log.Print(err)

	log.Print("start json parse")
	myLineRequest, err := UnmarshalLineRequest([]byte(request.Body))
	if err != nil {
		log.Fatal(err)
	}

	if len(myLineRequest.Events) == 0 {
		return events.APIGatewayProxyResponse{
			Body:       "aaa",
			StatusCode: 200,
		}, nil
	}

	log.Print("start create reply message")
	message := strings.Split(myLineRequest.Events[0].Message.Text, " ")
	source := myLineRequest.Events[0].Source
	var result = ""

	if message[0] == "bot" {
		switch message[1] {
		case "rand":
			result = RandomChoice(message[2:])

		case "id":
			result = source.GetId()

		default:
			result = "Please enter rand or id"
		}
		} else {
	}
	

	if _, err = bot.ReplyMessage(myLineRequest.Events[0].ReplyToken, linebot.NewTextMessage(result)).Do(); err != nil {
		log.Fatal(err)
	}
	log.Print(myLineRequest)
	log.Print(err)

	return events.APIGatewayProxyResponse{
		Body:       "aaa",
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(Handler)
}
