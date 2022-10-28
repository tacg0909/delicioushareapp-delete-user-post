package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func deleteDetail(postId string) error {
    err := deleteS3Object(fmt.Sprintf("large/%s.jpg", postId))
    if err != nil {
        return err
    }
    err = deleteDynamoDetailItem(postId)
    return err
}

func deleteDynamoDetailItem(postId string) error {
    db := dynamo.New(session.New(), &aws.Config{
        Region: aws.String(os.Getenv("DYNAMO_REGION")),
    })
    table := db.Table(os.Getenv("DETAIL_TABLE_NAME"))

    err := table.Delete("PostId", postId).Run()
    return err
}
