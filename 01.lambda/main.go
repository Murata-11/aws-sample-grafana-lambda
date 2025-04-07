package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type response struct {
	Datas []Data `json:"datas"`
}

type Data struct {
	Timestamp string `json:"timestamp"`
	Value     float64 `json:"value"`
}

func handleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (response, error) {
	return response{
		Datas: []Data{
			{
				Timestamp: "2025-04-07T00:00:00+09:00",
				Value:     10.10,
			},
			{
				Timestamp: "2025-04-07T00:30:00+09:00",
				Value:     20.20,
			},
			{
				Timestamp: "2025-04-07T01:00:00+09:00",
				Value:     30.30,
			},
		},
	}, nil

	// return response{}, nil
}

func main() {
	lambda.Start(handleRequest)
}