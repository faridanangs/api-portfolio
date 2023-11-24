package app

import (
	"database/sql"
	"fmt"
	"os"
	"rest_api_portfolio/helper"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func DatabaseConnect() *sql.DB {
	err := godotenv.Load(".env")
	helper.PanicIfError(err, "Error  godotenv.Load on app at DatabaseConnect")
	dbName := os.Getenv("DATABASE_NAME")
	user := os.Getenv("DATABASE_USER")
	localhost := os.Getenv("DATABASE_LOCALHOST")
	port := os.Getenv("DATABASE_PORT")

	SCRIPT := fmt.Sprintf("%s@tcp(%s:%s)/%s", user, localhost, port, dbName)
	db, err := sql.Open("mysql", SCRIPT)
	helper.PanicIfError(err, "Error sql.Open on app at DatabaseConnect")
	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(50)

	return db
}
