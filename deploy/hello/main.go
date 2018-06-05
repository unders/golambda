package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/unders/golambda/lambda/hello"
)

// These are set from the build script
var (
	Version    = "Version"
	Buildstamp = "Buildstamp"
	Githash    = "Githash"
)

func main() {
	l := hello.Lambda{Version: Version, BuildStamp: Buildstamp, Githash: Githash}
	lambda.Start(l.Handler)
}
