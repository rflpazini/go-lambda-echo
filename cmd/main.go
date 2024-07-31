package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rflpazini/articles/lambda/pkg/api/healthcheck"
	"github.com/rflpazini/articles/lambda/pkg/utils"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/healthcheck", healthcheck.Handler)
	e.GET("/healthcheck/info", healthcheck.InfoHandler)

	server := wrapRouter(e)
	lambda.Start(server)
}

func wrapRouter(e *echo.Echo) func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		body := strings.NewReader(request.Body)
		req := httptest.NewRequest(request.HTTPMethod, request.Path, body)
		for k, v := range request.Headers {
			req.Header.Add(k, v)
		}

		q := req.URL.Query()
		for k, v := range request.QueryStringParameters {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()

		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		res := rec.Result()
		responseBody, err := io.ReadAll(res.Body)
		if err != nil {
			return utils.FormatAPIErrorResponse(http.StatusInternalServerError, res.Header, err.Error())
		}

		return utils.FormatAPIResponse(res.StatusCode, res.Header, string(responseBody))
	}
}
