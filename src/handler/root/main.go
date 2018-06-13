package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	msg := fmt.Sprintf("Processing request data for request %s.\n", req.RequestContext.RequestID)
	msg += fmt.Sprintf("Body size = %d.\n", len(req.Body))
	msg += fmt.Sprintln("Headers:")
	for key, value := range req.Headers {
		msg += fmt.Sprintf("    %s: %s\n", key, value)
	}
	return events.APIGatewayProxyResponse{
		IsBase64Encoded: false,
		Headers:         map[string]string{"myheader": "myheader"},
		Body:            msg,
		StatusCode:      200,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
