rm -rf .aws-sam || true
rm api || true
GOARCH=amd64 GOOS=linux go build -o api ./service/api
sam local start-api
