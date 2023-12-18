# Lambda Example
This is a project for testing out lambda functions with go.

## Commands
```bash
aws iam create-role --role-name lambda-ex --assume-role-policy-document file://trust-policy.json
```

```bash
aws iam attach-role-policy --role-name lambda-ex --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
```

```bash
aws lambda create-function --function-name go-lambda-ex \
--zip-file fileb://function.zip --handler main --runtime go1.x \
--role arn:aws:iam::#:role/lambda-example
```

```bash
aws lambda invoke --function-name lambda-example --cli-binary-format raw-in-base64-out --payload '{"What is your name?": "Jim", "How old are you?", 33}' output.txt
```

```shell
zip function.zip main.exe
```