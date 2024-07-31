package utils

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func FormatAPIResponse(statusCode int, headers http.Header, responseData string) (events.APIGatewayProxyResponse, error) {
	responseHeaders := make(map[string]string)

	responseHeaders["Content-Type"] = "application/json"
	for key, value := range headers {
		responseHeaders[key] = ""

		if len(value) > 0 {
			responseHeaders[key] = value[0]
		}
	}

	responseHeaders["Access-Control-Allow-Origin"] = "*"
	responseHeaders["Access-Control-Allow-Headers"] = "origin,Accept,Authorization,Content-Type"

	return events.APIGatewayProxyResponse{
		Body:       responseData,
		Headers:    responseHeaders,
		StatusCode: statusCode,
	}, nil
}

func FormatAPIErrorResponse(statusCode int, headers http.Header, err string) (events.APIGatewayProxyResponse, error) {
	responseHeaders := make(map[string]string)

	responseHeaders["Content-Type"] = "application/json"
	for key, value := range headers {
		responseHeaders[key] = ""

		if len(value) > 0 {
			responseHeaders[key] = value[0]
		}
	}

	responseHeaders["Access-Control-Allow-Origin"] = "*"
	responseHeaders["Access-Control-Allow-Headers"] = "origin,Accept,Authorization,Content-Type"

	return events.APIGatewayProxyResponse{
		Body:       err,
		Headers:    responseHeaders,
		StatusCode: statusCode,
	}, nil
}
