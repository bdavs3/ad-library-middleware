package google

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

type Connection struct {
	ctx    context.Context
	client *bigquery.Client
}

func NewConnection(project string) (*Connection, error) {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, project)
	if err != nil {
		return nil, err
	}

	return &Connection{
		ctx:    ctx,
		client: client,
	}, nil
}

func (conn *Connection) RunQuery(body string) error {
	it, err := conn.client.Query(body).Read(conn.ctx)
	if err != nil {
		return err
	}

	// TODO: Print any output to an audit log
	for {
		var values []bigquery.Value
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(values)
	}

	return nil
}
