package score

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func Db() *dynamodb.DynamoDB {
	endpoint := "http://localhost:8000"
	return dynamodb.New(session.New(), &aws.Config{Region: aws.String("us-west-2"), Endpoint: &endpoint})
}

func DeleteTable() {
	db := Db()

	params := &dynamodb.DeleteTableInput{
		TableName: aws.String("Score"),
	}

	resp, err := db.DeleteTable(params)
	fmt.Println(resp, err)
}

func CreateTable() {
	db := Db()

	params := &dynamodb.CreateTableInput{
		TableName: aws.String("Score"),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Initials"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("Score"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Initials"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("Score"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
	}

	resp, err := db.CreateTable(params)
	fmt.Println(resp, err)
}

func Seed(initials string, score int) {
	db := Db()

	params := &dynamodb.PutItemInput{
		TableName: aws.String("Score"),
		Item: map[string]*dynamodb.AttributeValue{
			"Initials": {
				S: aws.String(initials),
			},
			"Score": {
				N: aws.String(fmt.Sprintf("%d", score)),
			},
		},
	}

	resp, err := db.PutItem(params)
	fmt.Println(resp, err)
}

func GetAll() (results interface{}, err error) {
	db := Db()

	params := &dynamodb.ScanInput{
		TableName: aws.String("Score"),
	}

	var resp *dynamodb.ScanOutput

	resp, err = db.Scan(params)
	results = resp.Items
	return
}
