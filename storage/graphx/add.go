package graphx

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"webcrawler/site"
)

func (g *Graph) AddWebsite(ctx context.Context, web site.Website) error {
	ses := g.Neo4j.NewSession(ctx, neo4j.SessionConfig{})
	_, err := ses.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `
				MERGE (w:Website {url: $url, ProminenceValue: $prominenceValue})
				on create set w.prominence = 1
on create set w.crawled_at = datetime()
				on match set w.prominence = coalesce(w.prominence, 0) + 1
				
`, map[string]interface{}{
			"url":             web.Url,
			"prominenceValue": web.ProminenceValue,
		})

		return nil, err
	})
	return err
}

func (g *Graph) AddLink(ctx context.Context, page site.Page) error {
	ses := g.Neo4j.NewSession(ctx, neo4j.SessionConfig{})
	var err error
	for _, link := range page.Links {
		_, err = ses.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
			_, err := tx.Run(ctx, `
				MERGE (page1:Page {url: $callURL})
				MERGE (page2:Page {url: $linkTo})
				MERGE (page1)-[:LINK_TO]->(page2)
				MERGE (page2)-[:LINK_FROM]->(page1)
				on create set page2.prominence = 1
				on create set page1.prominence = 1
				on match set page2.prominence = coalesce(page2.prominence, 0) + 1
				on create set page1.crawled_at = datetime()
				`, map[string]interface{}{
				"callURL": page.Url,
				"linkTo":  link,
			})
			return nil, err
		})

		if err != nil {
			return err
		}
	}
	return err
}

func (g *Graph) AddPage(ctx context.Context, web site.Website, page site.Page) error {
	ses := g.Neo4j.NewSession(ctx, neo4j.SessionConfig{})
	_, err := ses.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `
MERGE (page:Page {url:$url, title:$title})
MERGE (web:Website {baseURL:$baseURL})
MERGE (web)-[:OWNED_BY]->(page)
MERGE (page)-[:BELONGS_TO]->(web)
		`, map[string]any{
			"url":     page.Url,
			"title":   page.Title,
			"baseURL": web.Url,
		})

		return nil, err
	})
	return err
}
