package app

import (
	"context"
	"fmt"
	"time"

	"github.com/ChanKachan/business-card-product/internal/config"
	"github.com/ChanKachan/business-card-product/pkg"
	postgre "github.com/ChanKachan/business-card-product/pkg/postgreWrapper"
)

func Start() error {
	cfgPostgre := &config.CfgPostgre{}
	cfgPostgresMigrate := &config.ConfigPostgresMigration{}

	err := pkg.GetConfigsPath( /*"./", ".env",*/ []any{
		cfgPostgre,
		cfgPostgresMigrate,
	})
	if err != nil {
		return fmt.Errorf("start app: %v", err)
	}

	postgres := postgre.New(
		&postgre.Config{
			Host:                 cfgPostgre.Host,
			Port:                 cfgPostgre.Port,
			User:                 cfgPostgre.User,
			Password:             cfgPostgre.Password,
			DBName:               cfgPostgre.DB,
			SSLMode:              cfgPostgre.SSLMode,
			PostgresQueryTimeout: cfgPostgre.QueryTimeout,
		},
	)

	defer postgres.Close()

	err = postgres.NewPoolConfig(
		cfgPostgre.PoolConnectMax,
		time.Duration(cfgPostgre.PoolConnectIdleTime)*time.Second,
		time.Duration(cfgPostgre.PoolConnectLifeTime)*time.Second,
	)
	if err != nil {
		return fmt.Errorf("postgres new pool config: %v", err)
	}

	сtxConn, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// пулл подключений Postgres
	if err = postgres.ConnectionPool(сtxConn); err != nil {
		return fmt.Errorf("postgres connection pool: %v", err)
	}

	//postgresTx := postgres.GetPostgreTx()

	return nil
}
