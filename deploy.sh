aws lambda create-function \
    --function-name gotest \
    --zip-file fileb://function.zip \
    --handler main \
    --runtime go1.x \
    --role arn:aws:iam::659289708175:role/lambda-ex \
    --timeout 10
