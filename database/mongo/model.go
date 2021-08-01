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

func (model *Base) Database() string {
	return model.database
}

func (model *Base) Collection() string {
	return model.collection
}

func (model *Base) Client() Client {
	return model.client
}

func (model *Base) Do(ctx context.Context, exec Exec) (interface{}, error) {
	return model.Client().Do(ctx, model, exec)
}

func (model *Base) DoWithTransaction(ctx context.Context, exec Exec) (interface{}, error) {
	return model.Client().DoWithTransaction(ctx, model, exec)
}
