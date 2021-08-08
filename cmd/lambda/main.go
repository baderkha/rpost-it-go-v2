package main

import (
	"context"
	"rpost-it-go/internal/api"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadaptor "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
)

func main() {
	app, db := api.SetupDbAndRouter()
	adapter := fiberadaptor.New(app)
	lambda.Start(func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		// If no name is provided in the HTTP request body, throw an error
		return adapter.ProxyWithContext(ctx, req)
	})

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
