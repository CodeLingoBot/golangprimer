package main

/*
	go get -u github.com/aws/aws-sdk-go

	https://github.com/awsdocs/aws-doc-sdk-examples/tree/master/go/example_code
*/

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func main() {
	// Initialize a session in eu-west-1 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.

	aws_session := session.Must(session.NewSession(&aws.Config{Region: aws.String("eu-west-1")}))

	//sess := session.Must(session.NewSessionWithOptions(session.Options{
	//	Config: aws.Config{
	//		Region:           aws.String("eu-west-1"),
	//		EndpointResolver: endpoints.ResolverFunc(s3CustResolverFn),
	//	},
	//}))

	result := list_buckets(aws_session)

	print_aws_s3_bucket_list(result)
	}

func list_buckets(aws_sess *session.Session) *s3.ListBucketsOutput {
	// Create S3 service client

	//var svc *s3.S3 = s3.New(aws_sess)
	svc := s3.New(aws_sess)

	input := &s3.ListBucketsInput{}

	result, err := svc.ListBuckets(input)
	if err != nil {
		exitErrorf("Unable to list buckets, %v", err)
	}
	return result
}

func print_aws_s3_bucket_list(result *s3.ListBucketsOutput) {
	fmt.Println("Buckets:")
	for _, bucket := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(bucket.Name),
			aws.TimeValue(bucket.CreationDate))
	}
}

