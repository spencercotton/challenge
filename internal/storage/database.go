package storage

import (
	"database/sql/driver"
	"github.com/go-sql-driver/mysql"
)

// sqlConfig returns a mysql.Config for the given DSN.
func sqlConfig(dsn string) (*mysql.Config, error) {
	cfg, err := mysql.ParseDSN(dsn)
	if err != nil {
		return nil, err
	}
	cfg.MultiStatements = true
	return cfg, nil
}

// SQLConnector returns a driver.Connector for the given DSN.
func SQLConnector(dsn string) (driver.Connector, error) {
	cfg, err := sqlConfig(dsn)
	if err != nil {
		return nil, err
	}
	sqlConnector, err := mysql.NewConnector(cfg)
	if err != nil {
		return nil, err
	}
	return sqlConnector, nil
}
