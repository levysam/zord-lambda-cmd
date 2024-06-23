package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"go-skeleton/cmd"
	apigwhandler "go-skeleton/cmd/aws_lambda/apigw_handler"
	"go-skeleton/cmd/http/server"
)

func main() {
	cmd.Setup()
	echoserver := server.NewServer(cmd.Reg, cmd.ApiPrefix)
	handler := apigwhandler.NewAPIGWHandler(echoserver.Setup())
	lambda.Start(handler.Handler)
}
