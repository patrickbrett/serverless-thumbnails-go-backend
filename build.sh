GOARCH=amd64 GOOS=linux go build -o api ./service/api
sam local start-api