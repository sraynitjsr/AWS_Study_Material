package main

import (
    "context"
    "fmt"
  
    "github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
    Name string `json:"name"`
}

func Handler(ctx context.Context, event Event) (string, error) {
    return fmt.Sprintf("Hello, %s!", event.Name), nil
}

func main() {
    lambda.Start(Handler)
}
