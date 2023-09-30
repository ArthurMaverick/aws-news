package clients

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type ServiceSSM struct {
	SSM *ssm.Client
}

func NewServiceSSM() *ServiceSSM {
	cfg, err := NewClient().Client()
	if err != nil {
		fmt.Println(err.Error())
	}
	return &ServiceSSM{SSM: ssm.NewFromConfig(*cfg)}
}

func (s *ServiceSSM) PutParameter(value, KeyName string) (*ssm.PutParameterOutput, error) {

	putParameterInput := &ssm.PutParameterInput{
		Name:      aws.String(KeyName),
		Value:     aws.String(value),
		DataType:  aws.String("String"),
		Overwrite: aws.Bool(true),
	}
	payload, err := s.SSM.PutParameter(context.TODO(), putParameterInput)
	if err != nil {
		return nil, err
	}
	return payload, nil
}
