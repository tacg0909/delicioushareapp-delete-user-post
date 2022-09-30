package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func main() {
    lambda.Start(MeshiteroDeletePost)
}

type SearchQuery struct {
    UserId string `json:"userId"`
    PostedTime string `json:"postedTime"`
}

func MeshiteroDeletePost(c context.Context, searchQuery SearchQuery) error {
    db := dynamo.New(session.New(), &aws.Config{
        Region: aws.String("ap-northeast-1"),
    })
    table := db.Table("MeshiteroPostsTable")

    err := table.Delete("UserId", searchQuery.UserId).Range("PostedTime", searchQuery.PostedTime).Run()
    return err
}
