package person

import (
	"context"
	"log"

	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/types"
	"github.com/graniticio/granitic/v2/ws"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type AddPersonLogic struct {
	Log logging.Logger
}

type AddPersonRequest struct {
	Name   *types.NilableString `json:"name"`
	Age    *types.NilableInt64  `json:"age"`
	Gender *types.NilableString `json:"gender"`
}

type user struct {
	Name   string
	Age    int64  `dynamo:"Age"`
	Gender string `dynamo:"Gender"`
}

func (pr *AddPersonLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, cd *AddPersonRequest) {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("ap-south-1"),
		Endpoint: aws.String("http://localhost:8000")})
	if err != nil {
		log.Println(err)
	}
	db := dynamo.New(sess)
	table := db.Table("Users")
	u := user{Name: cd.Name.String(), Age: cd.Age.Int64(), Gender: cd.Gender.String()}
	errors := table.Put(u).Run()
	if errors != nil {
		log.Println(err)
	}
	pr.Log.LogInfof("New Person: '%s'", cd.Name.String())
	res.Body = CreatedResource{
		Name:   cd.Name.String(),
		Age:    cd.Age.Int64(),
		Gender: cd.Gender.String(),
	}
}

type CreatedResource struct {
	Name   string
	Age    int64
	Gender string
}
