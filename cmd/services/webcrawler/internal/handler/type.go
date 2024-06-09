package handler

import (
	"context"
	"webcrawler/cmd/services/webcrawler/internal/config"
	"webcrawler/cmd/services/webcrawler/internal/dynamoDBx"
	"webcrawler/cmd/services/webcrawler/internal/graphx"
	"webcrawler/pkg/queue"
	"webcrawler/pkg/site"
)

type Server struct {
	Queue  *queue.Handler
	Db     *dynamoDBx.DB
	Graph  *graphx.Graph
	Config *config.IgnoreList
}

type DBi interface {
	FetchWebsite(context.Context, string) (*site.Website, error)
	FetchPage(context.Context, string) (*site.Page, error)
	AddPage(context.Context, site.Page) error
	AddWebsite(context.Context, site.Website) error
	RemoveWebsite(context.Context, site.Website) error
	RemovePage(context.Context, site.Page) error
	UpdateWebsite(context.Context, site.Page, site.Website) error
}

func New(db *dynamoDBx.DB, queue *queue.Handler, graph *graphx.Graph, config *config.IgnoreList) Server {
	return Server{
		Db:     db,
		Queue:  queue,
		Graph:  graph,
		Config: config,
	}

}
