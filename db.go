package kweb

import (
	"github.com/jmoiron/sqlx"
	"github.com/kooksee/go-assert"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func (t *app) InitDb(cfg *DbConfig) {
	db, err := sqlx.Connect(cfg.Schema, cfg.DbUrl)
	assert.Err(err, "init db %s error")
	db.SetMaxIdleConns(10)
	t.dbs[cfg.Name] = db
}
