package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func deleteOutline(postId string, userId string, postedTime string) error {
    err := deleteS3Object(fmt.Sprintf("small/%s.jpg", postId))
    if err != nil {
        return err
    }
    err = deleteDynamoOutlineItem(postId, userId, postedTime)
    return err
}

func deleteDynamoOutlineItem(postId string, userId string, postedTime string) error {
    db := dynamo.New(session.New(), &aws.Config{
        Region: aws.String(os.Getenv("DYNAMO_REGION")),
    })
    table := db.Table(os.Getenv("OUTLINE_TABLE_NAME"))

    err := table.Delete("UserId", userId).Range("PostedTime", postedTime).Run()
    return err
}
