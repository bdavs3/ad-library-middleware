package database

import (
	"context"

	"cloud.google.com/go/bigquery"
)

type Connection struct {
	ctx     context.Context
	dataset *bigquery.Dataset
}

func NewConnection(project, dataset string) (*Connection, error) {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, project)
	ds := client.Dataset(dataset)

	if err != nil {
		return nil, err
	}

	return &Connection{
		ctx:     ctx,
		dataset: ds,
	}, nil
}

func (conn *Connection) Insert(table string, data interface{}) error {
	tb := conn.dataset.Table(table)
	u := tb.Inserter()

	if err := u.Put(conn.ctx, data); err != nil {
		return err
	}

	return nil
}

func (conn *Connection) Select(table string) *bigquery.RowIterator {
	tb := conn.dataset.Table(table)
	return tb.Read(conn.ctx)
}
