package student

import (
	"context"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
)

func NewLocalStackConfig(ctx context.Context) aws.Config {
	cfg, _ := config.LoadDefaultConfig(ctx, config.WithRegion("ap-northeast-1"))
	return cfg
}

var _ Database = (*DynamoDBDatabase)(nil)

type DynamoDBDatabase struct {
	tableName string
	client    *dynamodb.Client
}

type resolver struct{}

func (r *resolver) ResolveEndpoint(ctx context.Context, params dynamodb.EndpointParameters) (smithyendpoints.Endpoint, error) {
	u, err := url.Parse("http://localhost:4566")
	if err != nil {
		return smithyendpoints.Endpoint{}, err
	}

	return smithyendpoints.Endpoint{
		URI: *u,
	}, nil
}

func NewDynamoDBDatabase(ctx context.Context) *DynamoDBDatabase {
	cfg := NewLocalStackConfig(ctx)
	return &DynamoDBDatabase{
		tableName: "students",
		client:    dynamodb.NewFromConfig(cfg, dynamodb.WithEndpointResolverV2(&resolver{})),
	}
}

func (d *DynamoDBDatabase) Append(ctx context.Context, event IEvent) error {
	jsonStr, err := event.Json()
	if err != nil {
		return err
	}
	_, err = d.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(d.tableName),
		Item: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{
				Value: event.StreamId().String(),
			},
			"sk": &types.AttributeValueMemberS{
				Value: event.Sk(),
			},
			"event": &types.AttributeValueMemberS{
				Value: jsonStr,
			},
		},
	})
	return err
}

func (d *DynamoDBDatabase) GetStudent(studentId StudentId) *Student {
	// TBI
	return nil
}

func (d *DynamoDBDatabase) GetStudentView(studentId StudentId) *Student {
	// TBI
	return nil
}
