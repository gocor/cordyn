package cordyn

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// GetByCompositeKeyInput ...
type GetByCompositeKeyInput struct {
	DB             *dynamodb.DynamoDB
	TableName      string
	Key            CompositeKey
	ConsistentRead bool
}

// GetBySecondaryCompositeKeyInput ...
type GetBySecondaryCompositeKeyInput struct {
	DB             *dynamodb.DynamoDB
	TableName      string
	IndexName      string
	ConsistentRead bool
	Key            CompositeKey
}

// GetByAlternateCompositeKeyInput ...
type GetByAlternateCompositeKeyInput struct {
	DB             *dynamodb.DynamoDB
	TableName      string
	IndexName      string
	ConsistentRead bool
	Key            CompositeKey
	PKName         string
	SKName         string
}

// GetByKeyConditionInput ...
type GetByKeyConditionInput struct {
	DB               *dynamodb.DynamoDB
	TableName        string
	IndexName        string
	Expression       string
	ExpressionValues map[string]*dynamodb.AttributeValue
	ConsistentRead   bool
}

// PutInput ...
type PutInput struct {
	DB        *dynamodb.DynamoDB
	TableName string
	Item      interface{}
}
