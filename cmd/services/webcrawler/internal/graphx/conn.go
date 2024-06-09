package graphx

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func Conn(ctx context.Context, user string, password string, url string) (neo4j.DriverWithContext, error) {
	dbUri := url
	dbUser := user
	dbPassword := password
	driver, err := neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""),
	)

	if err != nil {
		return nil, err
	}
	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	}
	return driver, err
}

func New(graph neo4j.DriverWithContext) *Graph {
	return &Graph{
		Neo4j: graph,
	}

}
