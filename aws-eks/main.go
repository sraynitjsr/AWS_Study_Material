package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
)

func main() {
	svc := eks.New(session.New())
	input := &eks.CreateClusterInput{
		ClientRequestToken: aws.String("1d2129a1-3d38-460a-9756-e5b91fddb951"),
		Name:               aws.String("prod"),
		ResourcesVpcConfig: &eks.VpcConfigRequest{
			SecurityGroupIds: []*string{
				aws.String("sg-6979fe18"),
			},
			SubnetIds: []*string{
				aws.String("subnet-6782e71e"),
				aws.String("subnet-e7e761ac"),
			},
		},
		RoleArn: aws.String("arn:aws:iam::012345678910:role/eks-service-role-AWSServiceRoleForAmazonEKS-J7ONKE3BQ4PI"),
		Version: aws.String("1.10"),
	}

	result, err := svc.CreateCluster(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case eks.ErrCodeResourceInUseException:
				fmt.Println(eks.ErrCodeResourceInUseException, aerr.Error())
			case eks.ErrCodeResourceLimitExceededException:
				fmt.Println(eks.ErrCodeResourceLimitExceededException, aerr.Error())
			case eks.ErrCodeInvalidParameterException:
				fmt.Println(eks.ErrCodeInvalidParameterException, aerr.Error())
			case eks.ErrCodeClientException:
				fmt.Println(eks.ErrCodeClientException, aerr.Error())
			case eks.ErrCodeServerException:
				fmt.Println(eks.ErrCodeServerException, aerr.Error())
			case eks.ErrCodeServiceUnavailableException:
				fmt.Println(eks.ErrCodeServiceUnavailableException, aerr.Error())
			case eks.ErrCodeUnsupportedAvailabilityZoneException:
				fmt.Println(eks.ErrCodeUnsupportedAvailabilityZoneException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}
