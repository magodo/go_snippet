package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	_ "github.com/lib/pq"
)

type replStat struct {
	pid              int64
	usesysid         string
	usename          string
	application_name string
	client_addr      string
	client_hostname  string
	client_port      int64
	backend_start    time.Time
	backend_xmin     string
	state            string
	sent_location    string
	write_location   string
	flush_location   string
	replay_location  string
	sync_priority    int64
	sync_state       string
}

func main() {
	db, err := sql.Open("postgres",
		fmt.Sprintf(
			"user='%s' password='%s' host=%s port=%s dbname=postgres sslmode=disable connect_timeout=5",
			"postgres",
			"123",
			"172.255.255.254",
			"5432",
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	/*
		if err != nil {
			nerr := err.(*net.OpError)
			fmt.Println(nerr.Net)
		}
	*/
	if err != nil {
		log.Fatal(err)
	}

	/*
		res, err := Vquery(db, "SELECT * FROM pg_stat_replication where client_addr = $1", "172.20.0.3")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res[0]["state"])

		rows, err := db.Query("SELECT * FROM pg_stat_replication")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			stat := replStat{}
			err := rows.Scan(
				&stat.pid,
				&stat.usesysid,
				&stat.usename,
				&stat.application_name,
				&stat.client_addr,
				&stat.client_hostname,
				&stat.client_port,
				&stat.backend_start,
				&stat.backend_xmin,
				&stat.state,
				&stat.sent_location,
				&stat.write_location,
				&stat.flush_location,
				&stat.replay_location,
				&stat.sync_priority,
				&stat.sync_state,
			)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(stat)
		}

		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	*/
}
