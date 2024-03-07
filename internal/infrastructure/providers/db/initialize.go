package db

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gobuffalo/envy"
	"github.com/jmoiron/sqlx"
	"sync"
)

// структура соединения с базой
type DB struct {
	Conn  *sqlx.DB
	mutex sync.Mutex
}

func (d *DB) Initialize() *DB {
	host := envy.Get("DB_HOST", "")
	port := envy.Get("DB_PORT", "")
	name := envy.Get("DB_NAME", "")
	password := envy.Get("DB_PASS", "")
	user := envy.Get("DB_USER", ".env")

	url := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, password, host, port, name)

	conn, err := sqlx.Connect("mssql", url)
	if err != nil {
		panic(err)
	}

	d.Conn = conn

	return d
}

// Закрывает подключение к базе данных
func (h *DB) Close() error {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	if err := h.Conn.Close(); err != nil {
		return fmt.Errorf("closing db connection: %w", err)
	}

	return nil
}
