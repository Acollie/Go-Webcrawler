package queue

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func (h Handler) Fetch(ctx context.Context) ([]Message, error) {
	messages, err := h.client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(h.url),
		MaxNumberOfMessages: 10,
	})
	if err != nil {
		return nil, err
	}
	result, err := formatMessage(messages)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func formatMessage(messages *sqs.ReceiveMessageOutput) ([]Message, error) {
	var result []Message
	for _, msg := range messages.Messages {
		message := Message{}

		err := json.Unmarshal([]byte(*msg.Body), &message)
		if err != nil {
			return nil, err
		}
		result = append(result, Message{
			Url:     message.Url,
			Handler: msg.ReceiptHandle,
		})
	}
	return result, nil
}
