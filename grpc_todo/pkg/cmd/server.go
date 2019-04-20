package cmd

import (
	"context"
	"database/sql"
	"flag"
	"fmt"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/magodo/go_snippet/grpc_todo/pkg/protocol/grpc"
	"github.com/magodo/go_snippet/grpc_todo/pkg/protocol/rest"
	"github.com/magodo/go_snippet/grpc_todo/pkg/service"
)

type Config struct {
	GRPCPort string
	HTTPPort string

	DatastoreDBHost     string
	DatastoreDBUser     string
	DatastoreDBPassword string
	DatastoreDBSchema   string
}

func RunServer() error {
	ctx := context.Background()

	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.HTTPPort, "http-port", "", "http port to bind")
	flag.StringVar(&cfg.DatastoreDBHost, "db-host", "", "Database host")
	flag.StringVar(&cfg.DatastoreDBUser, "db-user", "", "Database user")
	flag.StringVar(&cfg.DatastoreDBPassword, "db-password", "", "Database password")
	flag.StringVar(&cfg.DatastoreDBSchema, "db-schema", "", "Database schema")
	flag.Parse()

	param := "parseTime=true" // (mysql specific) parse DATE/DATETIME to time.Time
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		cfg.DatastoreDBUser,
		cfg.DatastoreDBPassword,
		cfg.DatastoreDBHost,
		cfg.DatastoreDBSchema,
		param,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	defer db.Close()

	v1API := service.NewToDoServiceServer(db)
	go func() {
		rest.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
	}()
	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
