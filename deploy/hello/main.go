package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/unders/golambda/func/hello"
)

func main() {
	lambda.Start(hello.Handler)
}
