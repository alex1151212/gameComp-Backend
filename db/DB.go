package db

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Instance *gorm.DB
)

func InitMySQL() {

	var (
		UserName     string = os.Getenv("DB_USERNAME")
		Password     string = os.Getenv("DB_PASSWORD")
		Addr         string = os.Getenv("DB_HOST")
		Port, _             = strconv.Atoi(os.Getenv("DB_PORT"))
		Database     string = os.Getenv("DB_DATABASE")
		MaxLifetime  int    = 10
		MaxOpenConns int    = 10
		MaxIdleConns int    = 10
	)

	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", UserName, Password, Addr, Port, Database)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	conn, err := gorm.Open(mysql.Open(addr), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println("🔴 Connection to MySQL failed:", err)
		return
	}

	db, err1 := conn.DB()
	if err1 != nil {
		fmt.Println("🔴 Get Database failed:", err)
		return
	}
	db.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	db.SetMaxIdleConns(MaxIdleConns)
	db.SetMaxOpenConns(MaxOpenConns)

	Instance = conn
	fmt.Println("🟢 DB Connection Init Success.")
}
