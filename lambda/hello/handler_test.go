package hello

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestLambda_Handler(t *testing.T) {
	lambda := Lambda{
		Version:    "1",
		BuildStamp: "Tis Jun 5",
		Githash:    "18b6d0e7f7",
	}

	tests := []struct {
		request events.APIGatewayProxyRequest
		want    string
		err     error
	}{
		{
			// Test that the handler responds with the correct response
			// when a valid name is provided in the HTTP body
			request: events.APIGatewayProxyRequest{Body: "Paul"},
			want:    "Hello Paul",
			err:     nil,
		},
		{
			// Test that the handler responds ErrNameNotProvided
			// when no name is provided in the HTTP body
			request: events.APIGatewayProxyRequest{Body: ""},
			want:    "",
			err:     errNameNotProvided,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			response, err := lambda.Handler(test.request)
			if test.err != err {
				t.Errorf("\nWant: %s\n Got: %s\n", test.err, err)
			}
			if test.want != response.Body {
				t.Errorf("\nWant: %s\n Got: %s\n", test.want, response.Body)
			}
		})
	}
}
