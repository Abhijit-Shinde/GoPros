package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(req events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Print("Hello Lambda")

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
	}

	return response, nil
}
