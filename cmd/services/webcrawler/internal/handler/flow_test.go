package handler

import (
	"context"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"webcrawler/cmd/services/webcrawler/internal/dynamoDBx"
	"webcrawler/pkg/awsx"
	"webcrawler/pkg/queue"
)

func setup(t *testing.T, ctx context.Context) Server {
	cfg, err := awsx.GetConfig(ctx)
	require.NoError(t, err)

	sqsClient := queue.New(
		os.Getenv("LINKS_QUEUE"),
		cfg,
	)
	dbClient := dynamoDBx.New(
		os.Getenv("DB_TABLE_PAGE"),
		os.Getenv("DB_TABLE_WEBSITE"),
		cfg,
	)
	return Server{
		Queue: sqsClient,
		Db:    dbClient,
	}

}
func teardown() {
	println("teardown")

}

func TestFlow(t *testing.T) {

}
