package main

import (
	
	"context"
	"os"
	"strings"
	"github.com/aws/aws-lambda-go/events"
	"github.com/luiscdev/gambit/awsgo"
	"github.com/luiscdev/gambit/bd"
	"github.com/luiscdev/gambit/handlers"

	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(RunLambda)
}

func RunLambda(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	awsgo.InicializoAws()

	if !validateParams() {
		panic("Error en los parametros. debe enviar 'SecretName' y 'UrlPrefix'")
	}

	var res *events.APIGatewayProxyResponse
	path := strings.Replace(request.RawPath, os.Getenv("UrlPrefix"), "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	header := request.Headers

	bd.ReadSecret()

	status, message := handlers.Handlers(path, method, body, headers, request)

	headersResp := map[string]string{
		"Content-Type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		StatusCode: status
		Body: string(message)
		Headers: headersResp
	}

	return res, nil
}

func validateParams() bool {
	_, getParam := os.LookupEnv("SecretName")
	if !getParam {
		return getParam
	}
	
	_, getParam = os.LookupEnv("UrlPrefix")
	if !getParam {
		return getParam
	}

	return getParam
}