package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

var conn *pgxpool.Pool

// database connection using pool concept
// make the server to connect a pool instead of the database directly
// pool makes the persistent connection with database
func Connect(connString string) error {
	var err error;
	conn, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		return err
	}
	return nil
}

// simmillar like in node we make the time retuning function in go
func GetTime(ctx *gin.Context) time.Time {
	var time time.Time;
	err := conn.QueryRow(ctx, "SELECT NOW() as now;").Scan(&time)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	return time;
}