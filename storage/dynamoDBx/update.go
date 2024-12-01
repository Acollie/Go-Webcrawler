package dynamoDBx

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"reflect"
	"webcrawler/site"
)

func (db *DB) UpdateWebsite(ctx context.Context, page site.Page, website site.Website) error {
	websiteDB, err := db.FetchWebsite(ctx, page.BaseURL)
	if err != nil {
		return err
	}
	if reflect.DeepEqual(websiteDB, &site.Website{}) && err == nil {
		println("Website not found")
		websiteDB = &website
	} else {
		website.Links = append(websiteDB.Links, page.Url)
	}

	websiteDB.ProminenceValue += 1
	websiteDB.Links = append(websiteDB.Links, page.Url)
	website = *websiteDB

	av, err := attributevalue.MarshalMap(website)
	if err != nil {
		return err
	}
	_, err = db.session.PutItem(ctx, &dynamodb.PutItemInput{
		Item:      av,
		TableName: &db.websiteNameTable,
	})
	return err
}
