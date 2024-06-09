package main

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"
	localConfig "webcrawler/cmd/services/webcrawler/internal/config"
	"webcrawler/cmd/services/webcrawler/internal/dynamoDBx"
	"webcrawler/cmd/services/webcrawler/internal/graphx"
	"webcrawler/cmd/services/webcrawler/internal/handler"
	"webcrawler/pkg/awsx"
	"webcrawler/pkg/queue"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	ctx := context.Background()
	if err != nil {
		log.Fatalf("Failed to load .env file with error: %v", err)
	}

	if err != nil {
		log.Fatalf("Cannot load the AWS config: %s", err)
	}

	cfg, err := awsx.GetConfig(ctx)
	if err != nil {
		log.Fatalf("Cannot load the AWS config: %s", err)
	}
	sqsClient := queue.New(
		os.Getenv("LINKS_QUEUE"),
		cfg,
	)
	dbClient := dynamoDBx.New(
		os.Getenv("DB_TABLE_PAGE"),
		os.Getenv("DB_TABLE_WEBSITE"),
		cfg,
	)
	graphConn, err := graphx.Conn(ctx, os.Getenv("NEO4J_USER"), os.Getenv("NEO4J_PASSWORD"), os.Getenv("NEO4J_URL"))
	if err != nil {
		log.Fatalf("Cannot connect to the graph database: %s", err)
	}
	graph := graphx.New(graphConn)
	config := localConfig.Fetch()

	server := handler.New(dbClient, sqsClient, graph, config)

	initialLinks := []string{
		"https://blog.alexcollie.com/",
		"https://alexcollie.com",

		"https://reddit.com",
		"https://reddit.com/r/golang",
		"https://reddit.com/r/technology",
		"https://reddit.com/r/python",
		"https://reddit.com/r/javascript",
		"https://reddit.com/r/rust",
		"https://reddit.com/r/java",

		"https://bbc.co.uk",
		"https://bbc.co.uk/news",
		"https://bbc.co.uk/sport",
		"https://bbc.co.uk/weather",
		"https://bbc.co.uk/food",

		"https://cnn.com",
		"https://cnn.com/world",
		"https://cnn.com/us",

		"https://news.ycombinator.com",
		"https://news.ycombinator.com/newest",

		"https://hackernews.com",

		"https://techcrunch.com",
		"https://techcrunch.com/startups",

		"https://waitbutwhy.com/",

		"https://arstechnica.com/",

		"https://www.wired.com/",

		"https://www.theverge.com/",
	}

	server.Queue.AddFromString(ctx, initialLinks)
	server.Scan(ctx)

}
