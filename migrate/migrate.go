package migrate

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// CreateCoreTable ...
func CreateCoreTable(db *dynamodb.DynamoDB, tableName, idxSecondary, pkName, skName, pk2Name, sk2Name string) error {
	_, err := db.CreateTable(&dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String(pkName),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String(skName),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String(pk2Name),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String(sk2Name),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String(pkName),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String(skName),
				KeyType:       aws.String("RANGE"),
			},
		},
		BillingMode: aws.String(dynamodb.BillingModeProvisioned),
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
		TableName: aws.String(tableName),
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String(idxSecondary),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String(pk2Name),
						KeyType:       aws.String("HASH"),
					},
					{
						AttributeName: aws.String(sk2Name),
						KeyType:       aws.String("RANGE"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("KEYS_ONLY"),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(5),
					WriteCapacityUnits: aws.Int64(5),
				},
			},
		},
	})
	return err
}
