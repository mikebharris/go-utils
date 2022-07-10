package aws_config

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type AwsConfig struct {
	region                 string
	dynamoEndpointOverride string
}

func (a *AwsConfig) New(region string, dynamoEndpointOverride string) aws.Config {
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))
	if err != nil {
		panic(err)
	}

	if len(dynamoEndpointOverride) > 0 {
		defaultEndpointResolver := cfg.EndpointResolverWithOptions
		cfg.EndpointResolverWithOptions = aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			if service == dynamodb.ServiceID && len(dynamoEndpointOverride) > 0 {
				return aws.Endpoint{URL: dynamoEndpointOverride}, nil
			}
			return defaultEndpointResolver.ResolveEndpoint(service, region)
		})
	}

	return cfg
}
