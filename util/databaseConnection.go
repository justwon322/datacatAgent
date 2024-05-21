package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"datacatAgent/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql" // MySQL 드라이버
	"github.com/joho/godotenv"
)

func init() {
	// .env 파일 로드
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func OpenDB() (*ent.Client, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	// MySQL 데이터베이스 연결
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	client := ent.NewClient(ent.Driver(dialect.Debug(db)))
	return client, nil
}

func MigrateDB(client *ent.Client) {
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
