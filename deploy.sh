export S3_BUCKET=aws-sam-cli-managed-default-samclisourcebucket-13e1sa9930l7j
export STACK_NAME=serverless-thumbnails-go-backend

sam package --template-file template.yaml --s3-bucket $S3_BUCKET --output-template-file packaged.yaml
sam deploy --stack-name $STACK_NAME --template-file packaged.yaml --capabilities CAPABILITY_IAM
