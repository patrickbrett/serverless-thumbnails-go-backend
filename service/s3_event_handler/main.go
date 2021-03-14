package main

import (
	"context"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"

	"github.com/BurntSushi/graphics-go/graphics"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func blur_portion(ctx context.Context, event events.S3Event) {
	srcBucket := event.Records[0].S3.Bucket.Name
	dstBucket := srcBucket

	srcKey := event.Records[0].S3.Object.Key
	dstKey := "blurred/output.jpg"

	// Blur centre
	left := 0
	right := 100
	top := 0
	bottom := 100

	blurAmount := 18.0

	// Init S3

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-2")},
	)

	client := s3.New(sess)

	// Download image from S3

	downloadImageResponse, err := client.GetObject(&s3.GetObjectInput{Bucket: aws.String(srcBucket), Key: aws.String(srcKey)})

	fmt.Print(downloadImageResponse)

	img1, _, err := image.Decode(downloadImageResponse.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Create new blank image in memory and draw the loaded image onto it
	dstImage := image.NewRGBA(img1.Bounds())
	draw.Draw(dstImage, img1.Bounds(), img1, image.Point{0, 0}, draw.Src)

	// Create blurred version of image separately in memory
	blurred := image.NewRGBA(img1.Bounds())
	graphics.Blur(blurred, img1, &graphics.BlurOptions{StdDev: blurAmount})

	// Cut the desired portion of the blurred image and paste it on top of the copy of the original
	blurBounds := image.Rect(left, top, right, bottom)
	draw.Draw(dstImage, blurBounds, blurred, image.Point{left, top}, draw.Src)

	// Create file to save
	out, err := os.Create("/tmp/output.jpg")
	if err != nil {
		fmt.Println(err)
	}

	// Save as JPEG file
	var opt jpeg.Options
	opt.Quality = 80
	jpeg.Encode(out, dstImage, &opt)

	outputFile, err := os.Open("/tmp/output.jpg")
	if err != nil {
		fmt.Println(err)
	}

	client.PutObject(&s3.PutObjectInput{Body: outputFile, Bucket: aws.String(dstBucket), Key: aws.String(dstKey)})
}

func HandleRequest(ctx context.Context, event events.S3Event) (interface{}, error) {
	blur_portion(ctx, event)

	return nil, nil
}

func main() {
	lambda.Start(HandleRequest)
}
