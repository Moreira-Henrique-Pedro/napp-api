package infra

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

func Connect() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", username, password, host, port, dbname)
	var err error
	Conn, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	} else {
		fmt.Println("Connected to database")
	}
}

func Disconnect() {
	if Conn != nil {
		Conn.Close(context.Background())
	}
}
