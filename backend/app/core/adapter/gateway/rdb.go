package gateway

import "github.com/morning-night-guild/platform/pkg/ent"

// RDB RDBクライアント.
type RDB struct {
	*ent.Client
}

// RDBFactory RDBクライアントのファクトリ.
type RDBFactory interface {
	Of(dsn string) (*RDB, error)
}
