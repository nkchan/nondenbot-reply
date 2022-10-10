package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

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

func RandomChoice(data []string) string {
	rand.Seed(time.Now().UnixNano())
	z := rand.Intn(len(data))
	return data[z]
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
	Type   string `json:"type"`
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
	var result = ""

	if message[0] == "bot" {
		if message[1] == "rand" {
			result = RandomChoice(message[2:])
			log.Print("Start rand func")
		} else if len(message) >= 2 {
			result = message[1]
			log.Print("Start Reply Func")
		} else {
		}
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
