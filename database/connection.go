package database

import (
	"context"

	"cloud.google.com/go/bigquery"
)

// Connection is used to interface with a BigQuery database.
type Connection struct {
	ctx     context.Context
	dataset *bigquery.Dataset
}

// NewConnection establishes a Connection to the given BigQuery dataset. The provided dataset
// must be a valid dataset within the given project.
func NewConnection(project, dataset string) (*Connection, error) {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, project)
	if err != nil {
		return nil, err
	}
	ds := client.Dataset(dataset)

	return &Connection{
		ctx:     ctx,
		dataset: ds,
	}, nil
}

// Insert uploads data to the given table in the Connection's dataset. The data parameter
// must have a valid BigQuery interpretation as defined by custom types. See
// https://pkg.go.dev/cloud.google.com/go/bigquery#Inserter.Put for more information.
func (conn *Connection) Insert(table string, data interface{}) error {
	tb := conn.dataset.Table(table)
	u := tb.Inserter()

	if err := u.Put(conn.ctx, data); err != nil {
		return err
	}

	return nil
}

// Select retrieves all data from the given table in the Connection's dataset.
func (conn *Connection) Select(table string) *bigquery.RowIterator {
	tb := conn.dataset.Table(table)
	return tb.Read(conn.ctx)
}
