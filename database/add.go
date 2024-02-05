package database

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"webcrawler/site"
)

func (db *DB) AddPage(website site.Page) error {
	ctx := context.Background()
	av, err := attributevalue.MarshalMap(website)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(db.pageNameTable),
	}
	_, err = db.session.PutItem(ctx, input)
	return err
}
func (db *DB) AddWebsite(website site.Website) error {
	ctx := context.Background()
	av, err := attributevalue.MarshalMap(website)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(db.websiteNameTable),
	}
	_, err = db.session.PutItem(ctx, input)
	return err
}
