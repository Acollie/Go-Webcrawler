package dynamoDBx

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"webcrawler/site"
)

func (db *DB) RemovePage(ctx context.Context, website site.Page) error {
	_, err := db.session.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &db.pageNameTable,
		Key: map[string]types.AttributeValue{
			"PageURL": &types.AttributeValueMemberS{Value: website.Url},
		},
	})
	return err
}

func (db *DB) RemoveWebsite(ctx context.Context, website site.Website) error {
	_, err := db.session.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &db.websiteNameTable,
		Key: map[string]types.AttributeValue{
			"BaseURL": &types.AttributeValueMemberS{Value: website.Url},
		},
	})
	return err
}
