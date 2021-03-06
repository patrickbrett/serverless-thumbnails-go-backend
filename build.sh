rm -rf .aws-sam
rm api
GOARCH=amd64 GOOS=linux go build -o api ./service/api
sam local start-api