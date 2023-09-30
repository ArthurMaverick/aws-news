package clients

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type ServiceS3 struct {
	S3 *s3.Client
}

func NewServiceS3() *ServiceS3 {
	cfg, err := NewClient().Client()
	if err != nil {
		fmt.Println(err.Error())
	}
	return &ServiceS3{S3: s3.NewFromConfig(*cfg)}
}

func (s *ServiceS3) GetBucket(bucketName string) (*s3.CreateBucketOutput, error) {
	_, err := s.S3.HeadBucket(context.TODO(), &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})

	if err == nil {
		return nil, fmt.Errorf("bucket %s already exists", bucketName)
	}

	createBucketInput := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	}
	payload, err := s.S3.CreateBucket(context.TODO(), createBucketInput)
	if err != nil {
		return nil, err
	}
	return payload, nil
}
