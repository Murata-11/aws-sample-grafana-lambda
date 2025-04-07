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
	Value1     float64 `json:"value1"`
	Value2     float64 `json:"value2"`
	Value3     float64 `json:"value3"`
	Value4     float64 `json:"value4"`
	Value5     float64 `json:"value5"`
}

func handleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (response, error) {
	return response{
		Datas: []Data{
			{
				Timestamp: "2025-04-07T00:00:00+09:00",
				Value1:     10.10,
				Value2:     20.20,
				Value3:     30.30,
				Value4:     40.50,
			},
		},
	}, nil

	// return response{}, nil
}

func main() {
	lambda.Start(handleRequest)
}