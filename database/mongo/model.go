package mongo

import "context"

type Model interface {
	Database() string
	Collection() string
}

type Base struct {
	database   string
	collection string
	ctx        context.Context
	client     Client
}

func NewBaseModel(client Client, database string, collection string) *Base {
	return &Base{client: client, database: database, collection: collection}
}
