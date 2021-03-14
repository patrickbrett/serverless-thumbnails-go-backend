package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

func putImage() interface{} {
	bucket := "thumbnails-go-angular"
	prefix := "full-size"

	filename := uuid.New().String() + ".jpg"
	s3Key := prefix + "/" + filename

	fmt.Print(filename)

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-2")},
	)

	client := s3.New(sess)

	req, _ := client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(s3Key),
	})
	urlStr, _ := req.Presign(15 * time.Minute)

	return map[string]interface{}{"uploadUrl": urlStr, "filename": filename}
}
