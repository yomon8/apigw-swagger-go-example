package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Message struct {
	Message string `json:"message"`
}

func getResponse(msg string, statusCode int) events.APIGatewayProxyResponse {
	json, _ := json.Marshal(Message{Message: msg})
	return events.APIGatewayProxyResponse{
		IsBase64Encoded: false,
		Headers:         map[string]string{"custom_header": "custom_header_value"},
		Body:            string(json),
		StatusCode:      statusCode,
	}
}

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if name := req.QueryStringParameters["name"]; name != "" {
		return getResponse(fmt.Sprintf("Hello, %s !", name), 200), nil
	}
	return getResponse("What`s your name?", 400), nil
}

func main() {
	lambda.Start(HandleRequest)
}
