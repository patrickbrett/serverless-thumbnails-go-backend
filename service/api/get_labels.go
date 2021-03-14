package main

import (
	"net/http"
)

type LabelRequest struct {
	Filename string
}

type Label struct {
	Top       int
	Left      int
	Bottom    int
	Right     int
	LabelType string
}

func getLabels(r *http.Request) interface{} {
	// bucket := "thumbnails-go-angular"
	// prefix := "full-size"

	// TODO: load filename from request QSP's

	// s3Key := prefix + "/" + body.Filename

	// sess, _ := session.NewSession(&aws.Config{
	// 	Region: aws.String("ap-southeast-2")},
	// )

	// client := s3.New(sess)

	var labels = []Label{
		{
			Top:       10,
			Left:      200,
			Bottom:    50,
			Right:     20,
			LabelType: "number_plate",
		},
	}

	return map[string]interface{}{"labels": labels}
}
