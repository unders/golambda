package hello

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/pkg/errors"
)

const lambda = "lambda=hello requestID=%s version=%s buildstamp=%s githash=%s"

var (
	// errNameNotProvided is thrown when a name is not provided
	errNameNotProvided = errors.New("no name was provided in the HTTP body")
)

// Lambda that implements hello handler
type Lambda struct {
	Version    string
	BuildStamp string
	Githash    string
}

// Handler a lambda function that says hello
func (l Lambda) Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf(lambda, request.RequestContext.RequestID, l.Version, l.BuildStamp, l.Githash)

	// If no name is provided in the HTTP request body, throw an error
	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, errNameNotProvided
	}

	return events.APIGatewayProxyResponse{
		Body:       "Hello " + request.Body + "\n",
		StatusCode: 200,
	}, nil
}
