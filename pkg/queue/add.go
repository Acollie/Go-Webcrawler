package queue

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"strconv"
)

func (h Handler) Add(ctx context.Context, message Message) error {
	messageJson, err := json.Marshal(message)
	if err != nil {
		return err
	}
	_, err = h.client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    aws.String(h.url),
		MessageBody: aws.String(string(messageJson)),
	})
	return err
}
func (h Handler) AddFromString(ctx context.Context, urls []string) error {
	var message []Message
	for _, msg := range urls {
		message = append(message, NewMessage(msg))
	}
	return h.BatchAdd(ctx, message)
}

func (h Handler) BatchAdd(ctx context.Context, messages []Message) error {
	const batchSize = 10
	for i := 0; i < len(messages); i += batchSize {
		end := i + batchSize
		if end > len(messages) {
			end = len(messages)
		}
		batch := messages[i:end]

		var entries []types.SendMessageBatchRequestEntry
		for j, msg := range batch {
			messageJson, err := json.Marshal(msg)
			if err != nil {
				return err
			}

			entry := types.SendMessageBatchRequestEntry{
				Id:          aws.String(strconv.Itoa(j)),
				MessageBody: aws.String(string(messageJson)),
			}

			entries = append(entries, entry)
		}

		_, err := h.client.SendMessageBatch(ctx, &sqs.SendMessageBatchInput{
			QueueUrl: aws.String(h.url),
			Entries:  entries,
		})

		if err != nil {
			return err
		}
	}
	return nil
}
