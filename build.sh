rm -rf .aws-sam || true
rm api || true
rm s3_event_handler || true
GOARCH=amd64 GOOS=linux go build -o api ./service/api
GOARCH=amd64 GOOS=linux go build -o s3_event_handler ./service/s3_event_handler
