aws dynamodb create-table \
    --table-name students \
    --attribute-definitions AttributeName=pk,AttributeType=S \
    AttributeName=sk,AttributeType=S \
    --key-schema AttributeName=pk,KeyType=HASH \
    AttributeName=sk,KeyType=RANGE \
    --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --endpoint-url http://localhost:4566
