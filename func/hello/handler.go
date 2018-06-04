package hello

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/pkg/errors"
)

var (
	// errNameNotProvided is thrown when a name is not provided
	errNameNotProvided = errors.New("no name was provided in the HTTP body")
)

// Handler a lambda function that says hello
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	// If no name is provided in the HTTP request body, throw an error
	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, errNameNotProvided
	}

	return events.APIGatewayProxyResponse{
		Body:       "Hello " + request.Body,
		StatusCode: 200,
	}, nil
}
