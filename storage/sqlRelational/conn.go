package sqlRelational

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type SqlDB struct {
	Client *sql.DB
	DBName string
}

func connect(host string, password string, user string, dbName string, port int) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName)
	client, err := sql.Open("mysql", dsn) // Use the MySQL driver
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	if err = client.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return client, nil
}

func New(host string, password string, user string, dbName string, port int) (*SqlDB, error) {
	client, err := connect(host, password, user, dbName, port)
	if err != nil {
		return nil, err
	}
	return &SqlDB{
		Client: client,
		DBName: dbName,
	}, nil
}
