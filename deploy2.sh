GOOS=linux go build
zip function.zip site_modify_checker
aws lambda update-function-code --function-name gotest --zip-file fileb://function.zip
