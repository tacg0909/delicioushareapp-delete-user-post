package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
    lambda.Start(delicioushareappDeleteUserPost)
}

type SearchQuery struct {
    PostId string `json:"postId"`
    UserId string `json:"userId"`
    PostedTime string `json:"postedTime"`
}

func delicioushareappDeleteUserPost(q SearchQuery) error {
    err := deleteDetail(q.PostId)
    if err != nil {
        return err
    }
    err = deleteOutline(q.PostId, q.UserId, q.PostedTime)
    return err
}
