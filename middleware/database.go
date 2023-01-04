package middleware

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
	Err    error
)

func PostgreSQLConnect() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading Env File: ", err)
	}

	environment := os.Getenv("ENVIRONMENT")
	fmt.Println(environment)
	err = godotenv.Load(fmt.Sprintf(".env-%s", environment))
	if err != nil {
		log.Fatal("Error Loading Env File: ", err)
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_UNAME")
	pass := os.Getenv("DB_PASS")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbname := os.Getenv("DB_NAME")

	// uri := os.Getenv("DB_URI")
	fmt.Println(user)
	dsn := fmt.Sprintf("host=%s user="+user+" password=%s dbname=%s port=%d sslmode=disable", host, pass, dbname, port)
	fmt.Println(dsn)
	DBConn, Err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})

	DBConn.AutoMigrate()
}
