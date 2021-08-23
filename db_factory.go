package cordyn

import (
	"sync"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var sharedDB *dynamodb.DynamoDB
var sharedDBOnce sync.Once

// NewDB ...
func NewDB(awsSession *session.Session) *dynamodb.DynamoDB {
	sharedDBOnce.Do(func() {
		sharedDB = dynamodb.New(awsSession)
	})
	return sharedDB
}
