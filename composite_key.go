package cordyn

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// CompositeKey ...
type CompositeKey struct {
	PK string
	SK string
}

// ToAttrValueMap converts the CompositeKey to a map[string]*Dynamodb.AttributeValue
func (k CompositeKey) ToAttrValueMap() map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		"pk": {S: aws.String(k.PK)},
		"sk": {S: aws.String(k.SK)},
	}
}

// ToAttrValueMapWithKeys converts the CompositeKey to a map[string]*Dynamodb.AttributeValue
func (k CompositeKey) ToAttrValueMapWithKeys(pkName, skName string) map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		pkName: {S: aws.String(k.PK)},
		skName: {S: aws.String(k.SK)},
	}
}
