package conf

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	err error
	db  *sql.DB
	DB  *gorm.DB
)

func InitDB() func() {
	DB, err = gorm.Open(postgres.Open(connectionOptions()), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic(err)
	}
	db, err = DB.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(30)
	db.SetMaxOpenConns(100)

	go func() {
		t := time.NewTimer(5 * time.Second)
		for {
			select {
			case <-t.C:
				if err := db.Ping(); err != nil {
					log.Printf("postgre connect error: %v", err)
				}
			}
		}
	}()

	return func() {
		_ = db.Close()
	}
}

func connectionOptions() string {
	c := Conf.Postgre
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", c.Username, c.Password, c.Databases, c.Host, c.Port)
}
