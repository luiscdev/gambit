package handlers

import (
	"fmt"
	// "strconv"

	"github.com/aws/aws-lambda-go/events"
)

func Handlers(path string, method string, body, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Voy a procesar ", path, ">", method)

	// id := request.PathParameters["id"]
	// idn, _ := strconv.Atoi(id)

	return 400, "method invalid"
}
