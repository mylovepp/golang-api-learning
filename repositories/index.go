package repositories

import (
	"context"
	conn "database/sql"
	"fmt"
	"time"

	"mylovepp/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	_ "github.com/go-sql-driver/mysql"
)

var entClient *ent.Client

func InitialDbConnection() (*ent.Client, error) {
	username := "mylovepp_admin"
	password := "pp43401439!Db"
	host := "db4free.net"
	port := "3306"
	dbName := "mylovepp_db"

	connDb, err := conn.Open(dialect.MySQL, username+":"+password+"@tcp("+host+":"+port+")/"+dbName+"?parseTime=true")

	if err != nil {
		panic(err.Error())
	}

	connDb.SetMaxIdleConns(10)
	connDb.SetMaxOpenConns(100)
	connDb.SetConnMaxLifetime(time.Hour)

	mySqlDriver := entsql.OpenDB(dialect.MySQL, connDb)
	entClient = ent.NewClient(ent.Driver(mySqlDriver))

	// if err := entClient.Schema.Create(context.Background()); err != nil {
	// 	return nil, err
	// }

	return entClient, nil
}

func withTransaction(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := entClient.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
