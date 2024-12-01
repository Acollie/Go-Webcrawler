package dynamoDBx

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"webcrawler/site"
)

func (db *DB) AddPage(ctx context.Context, page site.Page) error {
	av, err := attributevalue.MarshalMap(page)
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

func (db *DB) AddWebsite(ctx context.Context, website site.Website) error {
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
