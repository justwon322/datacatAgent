package db

/**
1. author : 유용원
2. 주요 기능 (순서대로)
2-1). '.env' 파일에 입력되어 있는 DB 설정 정보 바인딩
2-2). db 커넥션 맺음
**/
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
	err := godotenv.Load(".env")
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
