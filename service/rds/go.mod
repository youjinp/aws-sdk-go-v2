module github.com/aws/aws-sdk-go-v2/service/rds

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v0.29.0
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v0.1.1
	github.com/awslabs/smithy-go v0.3.0
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/service/internal/presigned-url => ../../service/internal/presigned-url/
