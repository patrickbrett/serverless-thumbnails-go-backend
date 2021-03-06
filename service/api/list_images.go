package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"fmt"
)

func listImages() []string {
	bucket := "npb-source-bucket-sourcebucket-96mim45irkcp"
	prefix := "images/train"
	keyLimit := int64(1000)

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-2")},
	)

	client := s3.New(sess)
	result, err := client.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket), Prefix: aws.String(prefix), MaxKeys: &keyLimit})
	if err != nil {
		panic(err)
	}

	keys := make([]string, *result.KeyCount)
	for i, s := range result.Contents {
		keys[i] = *s.Key
	}

	fmt.Print(keys)

	return keys
}
