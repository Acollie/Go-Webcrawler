package queue

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"os"
)

type HandlerI interface {
	Fetch(ctx context.Context) ([]Message, error)
	Add(ctx context.Context, message Message) error
	BatchAdd(ctx context.Context, message []Message) error
	Remove(ctx context.Context, message string) error
}
type Handler struct {
	url    string
	client *sqs.Client
}

type Message struct {
	Url     string  `json:"url"`
	Handler *string `json:"handler"`
}

func New(url string, cfg aws.Config) *Handler {
	if os.Getenv("ENVIRONMENT") == "local" {
		cfg.EndpointResolver = aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           "http://localhost:4570",
				SigningRegion: "us-west-2",
			}, nil
		})
	}
	return &Handler{
		url:    url,
		client: sqs.NewFromConfig(cfg),
	}
}

func NewMessage(url string) Message {
	return Message{
		Url: url,
	}
}
