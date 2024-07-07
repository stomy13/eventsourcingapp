package student

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
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
	eventJson, err := event.Json()
	if err != nil {
		return err
	}

	// TODO: Use Transaction
	// synchronous projection
	// update student view
	student, err := d.GetStudentView(ctx, event.StreamId())
	if err != nil {
		return err
	}
	// no student created
	if student == nil {
		k := fmt.Sprintf("%s%s", event.StreamId().String(), "_view")
		student = &Student{
			Pk: k,
			Sk: k,
		}
	}
	student.Apply(event)
	studentJson, err := student.Json()
	if err != nil {
		return err
	}

	_, err = d.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(d.tableName),
		Item: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{
				Value: student.Pk,
			},
			"sk": &types.AttributeValueMemberS{
				Value: student.Sk,
			},
			"student": &types.AttributeValueMemberS{
				Value: studentJson,
			},
		},
	})
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
				Value: eventJson,
			},
		},
	})
	return err
}

type DynamoDBObject struct {
	Event string
}

func (obj *DynamoDBObject) ToEvent() Event {
	event, _ := NewEventFromJson(obj.Event)
	return event
}

func (d *DynamoDBDatabase) GetStudent(ctx context.Context, studentId StudentId) (*Student, error) {
	output, err := d.client.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String(d.tableName),
		KeyConditionExpression: aws.String("pk = :pk"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{
				Value: studentId.String(),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	if len(output.Items) == 0 {
		return nil, nil
	}

	var events []IEvent
	for _, item := range output.Items {
		obj := &DynamoDBObject{}
		if err := attributevalue.UnmarshalMap(item, obj); err != nil {
			return nil, err
		}

		event, err := NewEventFromJson(obj.Event)
		if err != nil {
			return nil, err
		}

		switch event.Type {
		case "StudentCreated":
			e := &StudentCreated{}
			if err := json.Unmarshal([]byte(obj.Event), e); err != nil {
				return nil, err
			}
			events = append(events, e)
		case "StudentUpdated":
			e := &StudentUpdated{}
			if err := json.Unmarshal([]byte(obj.Event), e); err != nil {
				return nil, err
			}
			events = append(events, e)
		case "StudentEnrolled":
			e := &StudentEnrolled{}
			if err := json.Unmarshal([]byte(obj.Event), e); err != nil {
				return nil, err
			}
			events = append(events, e)
		}
	}
	if len(events) == 0 {
		return nil, errors.New("student not found")
	}

	student := &Student{}
	for _, event := range events {
		event.apply(student)
	}
	return student, nil
}

func (d *DynamoDBDatabase) GetStudentView(ctx context.Context, studentId StudentId) (*Student, error) {
	k := fmt.Sprintf("%s%s", studentId.String(), "_view")
	output, err := d.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(d.tableName),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{
				Value: k,
			},
			"sk": &types.AttributeValueMemberS{
				Value: k,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	studentAttr := output.Item["student"]
	if studentAttr == nil {
		return nil, nil
	}
	studentAttrS, ok := studentAttr.(*types.AttributeValueMemberS)
	if !ok {
		return nil, errors.New("invalid student attribute")
	}
	student, err := NewStudentFromJson(studentAttrS.Value)
	if err != nil {
		return nil, err
	}
	return student, nil
}
