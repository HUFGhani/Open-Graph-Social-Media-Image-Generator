package main

import (
	"context"
	"functions/awsS3"
	"functions/openGraph"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
)

var fiberLambda *fiberadapter.FiberLambda

func main() {
	lambda.Start(Handler)
}
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return fiberLambda.ProxyWithContext(ctx, req)
}

func init() {
	awsS3.S3downloadAssets()
	log.Printf("Fiber cold start")
	var app *fiber.App
	app = fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		pagetitle := c.Query("title")
		openGraph.CreateOpenGraphImage(pagetitle)
		return c.SendFile("/tmp/outputFilename.png")
		// return c.SendString(pagetitle)
	})
	fiberLambda = fiberadapter.New(app)
}
