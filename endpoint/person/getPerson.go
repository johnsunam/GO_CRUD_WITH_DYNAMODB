package person

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/graniticio/granitic/types"
	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/ws"
	// "github.com/guregu/dynamo"
)

type GetPersonLogic struct {
	EnvLabel string
	Log      logging.Logger
}

type Person struct {
	Name   string `json:"Name"`
	Age    int64  `json:"Age"`
	Gender string `json:"Gender"`
}

type User struct {
	Name   string `dynamodbav:"Name"`
	Age    int64  `dynamodbav:"Age"`
	Gender string `dynamodbav:"Gender"`
}

type UserQuery struct {
	NAME          string
	NormaliseName *types.NilableBool
}

func (gl *GetPersonLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, q *UserQuery) {
	sess, sessionErr := session.NewSession(&aws.Config{
		Region:   aws.String("ap-south-1"),
		Endpoint: aws.String("http://localhost:8000")})
	if sessionErr != nil {
		log.Println(sessionErr)
	}
	db := dynamodb.New(sess)
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Name": {
				S: aws.String(q.NAME),
			},
		},
		TableName: aws.String("Users"),
	}

	result, err := db.GetItem(input)
	if err != nil {
		log.Println(err.Error())

	}
	user := Person{}
	errs := dynamodbattribute.UnmarshalMap(result.Item, &user)

	if errs != nil {
		log.Println(errs.Error())

	}
	fmt.Println(result.Item)
	res.Body = user
	// fmt.Println(results)
}

type UserReturn struct {
	Name   string
	Age    int64
	Gender string
}
