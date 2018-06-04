package main_test

import (
	"testing"
	"github.com/unders/golambda/func/hello"

	"github.com/aws/aws-lambda-go/events"
	"reflect"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request events.APIGatewayProxyRequest
		want  string
		err     error
	}{
		{
			// Test that the handler responds with the correct response
			// when a valid name is provided in the HTTP body
			request: events.APIGatewayProxyRequest{Body: "Paul"},
			want:  "Hello Paul",
			err:     nil,
		},
		{
			// Test that the handler responds ErrNameNotProvided
			// when no name is provided in the HTTP body
			request: events.APIGatewayProxyRequest{Body: ""},
			want:  "",
			err:     main.ErrNameNotProvided,
		},
	}


	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			response, err := main.Handler(test.request)
			if test.err !=  err {
				t.Errorf("\nWant: %s\n Got: %s\n", test.err, err)
			}
			if !reflect.DeepEqual(test.want, response.Body) {
				t.Errorf("\nWant: %s\n Got: %s\n", test.want, response.Body)
			}
		})
	}
}
