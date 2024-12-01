package dynamoDBx

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"webcrawler/site"
)

func (db *DB) FetchWebsite(ctx context.Context, website string) (*site.Website, error) {
	resp, err := db.session.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &db.websiteNameTable,
		Key: map[string]types.AttributeValue{
			"BaseURL": &types.AttributeValueMemberS{Value: website},
		},
	})
	if err != nil {
		return &site.Website{}, err
	}
	if resp.Item == nil {
		return &site.Website{}, nil
	}
	websiteFormat, err := formatWebsite(resp)
	if err != nil {
		return &site.Website{}, err
	}
	return &websiteFormat, nil
}

func (db *DB) FetchPage(ctx context.Context, website string) (*site.Page, error) {
	resp, err := db.session.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &db.pageNameTable,
		Key: map[string]types.AttributeValue{
			"PageURL": &types.AttributeValueMemberS{Value: website},
		},
	})
	if err != nil {
		return nil, err
	}
	if resp != nil {
		return nil, nil
	}
	page, err := formatPage(ctx, resp)
	if err != nil {
		return nil, err
	}
	return &page, nil
}

func formatPage(ctx context.Context, input *dynamodb.GetItemOutput) (site.Page, error) {
	var page site.Page
	err := attributevalue.UnmarshalMap(input.Item, &page)
	if err != nil {
		return site.Page{}, err
	}
	return page, nil
}

func formatWebsite(input *dynamodb.GetItemOutput) (site.Website, error) {
	var website site.Website
	err := attributevalue.UnmarshalMap(input.Item, &website)
	if err != nil {
		return site.Website{}, err
	}
	return website, nil

}
