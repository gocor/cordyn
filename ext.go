package cordyn

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/pkg/errors"
)

// GetByCompositeKey ...
func GetByCompositeKey(ctx context.Context, result interface{}, input GetByCompositeKeyInput) error {
	getInput := dynamodb.GetItemInput{
		TableName:      aws.String(input.TableName),
		ConsistentRead: aws.Bool(input.ConsistentRead),
		Key:            input.Key.ToAttrValueMap(),
	}

	output, err := input.DB.GetItem(&getInput)
	if err != nil {
		return errors.WithStack(err)
	}
	if output == nil {
		return errors.New("Unexpected nil output")
	}

	// no result found
	if output.Item == nil || len(output.Item) == 0 {
		return nil
	}

	if err := dynamodbattribute.UnmarshalMap(output.Item, &result); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// GetByKeyCondition ...
func GetByKeyCondition(ctx context.Context, result interface{}, input GetByKeyConditionInput) error {
	queryInput := &dynamodb.QueryInput{
		TableName:      aws.String(input.TableName),
		IndexName:      aws.String(input.IndexName),
		ConsistentRead: aws.Bool(input.ConsistentRead),
		// Limit:                     aws.Int64(1),
		KeyConditionExpression:    aws.String(input.Expression),
		ExpressionAttributeValues: input.ExpressionValues,
	}

	output, err := input.DB.QueryWithContext(ctx, queryInput)
	return validateGetQueryOutput(ctx, output, err, result)
}

// GetBySecondayCompositeKey ...
func GetBySecondayCompositeKey(ctx context.Context, result interface{}, input GetBySecondaryCompositeKeyInput) error {
	alternateKeyInput := GetByAlternateCompositeKeyInput{
		DB:             input.DB,
		TableName:      input.TableName,
		IndexName:      input.IndexName,
		ConsistentRead: input.ConsistentRead,
		Key:            input.Key,
		PKName:         "pk2",
		SKName:         "sk2",
	}

	return GetByAlternateCompositeKey(ctx, result, alternateKeyInput)
}

// GetByAlternateCompositeKey ...
func GetByAlternateCompositeKey(ctx context.Context, result interface{}, input GetByAlternateCompositeKeyInput) error {
	keyExpr := fmt.Sprintf("%s = :pk AND %s = :sk", input.PKName, input.SKName)

	queryInput := &dynamodb.QueryInput{
		TableName:      aws.String(input.TableName),
		IndexName:      aws.String(input.IndexName),
		ConsistentRead: aws.Bool(input.ConsistentRead),
		// Limit:                     aws.Int64(1),
		KeyConditionExpression:    aws.String(keyExpr),
		ExpressionAttributeValues: input.Key.ToAttrValueMapWithKeys(":pk", ":sk"),
	}

	output, err := input.DB.QueryWithContext(ctx, queryInput)
	return validateGetQueryOutput(ctx, output, err, result)
}

func validateGetQueryOutput(ctx context.Context, output *dynamodb.QueryOutput, err error, result interface{}) error {
	if err != nil {
		return err
	}
	if output == nil {
		return errors.New("Unexpected nil output")
	}

	// no results found
	if output.Items == nil || output.Count == nil || *output.Count == int64(0) {
		return nil
	}
	if *output.Count > int64(1) {
		return errors.New(fmt.Sprintf("Too many items (%d) returned", *output.Count))
	}

	if err := dynamodbattribute.UnmarshalMap(output.Items[0], &result); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Put ...
func Put(ctx context.Context, input PutInput) error {
	item, err := dynamodbattribute.MarshalMap(input.Item)
	if err != nil {
		return errors.WithStack(err)
	}

	putInput := dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(input.TableName),
	}

	_, err = input.DB.PutItemWithContext(ctx, &putInput)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
