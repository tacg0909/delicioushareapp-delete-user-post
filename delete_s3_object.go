package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func deleteS3Object(key string) error {
    client := s3.New(
        session.New(),
        &aws.Config{ Region: aws.String(os.Getenv("S3_REGION")) },
    )
    _, err := client.DeleteObject(&s3.DeleteObjectInput{
        Bucket: aws.String(os.Getenv("BUCKET_NAME")),
        Key: aws.String(key),
    })
    return err
}
