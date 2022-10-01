package database

import (
	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/morning-night-guild/platform/app/core/adapter/gateway"
	"github.com/morning-night-guild/platform/pkg/ent"
)

var _ gateway.RDBFactory = (*Client)(nil)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c Client) Of(dsn string) (*gateway.RDB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return &gateway.RDB{}, err
	}

	drv := entsql.OpenDB(dialect.Postgres, db)

	client := ent.NewClient(ent.Driver(drv))

	return &gateway.RDB{
		Client: client,
	}, err
}
