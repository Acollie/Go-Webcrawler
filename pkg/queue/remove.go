package queue

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func (h Handler) Remove(ctx context.Context, receiptHandler string) error {
	_, err := h.client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(h.url),
		ReceiptHandle: aws.String(receiptHandler),
	})
	return err
}
