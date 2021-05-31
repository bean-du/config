package conf

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var (
	err error
	db  *sql.DB
	DB  *gorm.DB
)



func InitDB() func() {
	if Conf.DBType == "sqlite" {
		DB, err = gorm.Open(sqlite.Open(Conf.Sqlite), &gorm.Config{})
	}else {
		DB, err = gorm.Open(postgres.Open(connectionOptions()), &gorm.Config{
			PrepareStmt: true,
		})
	}
	if err != nil {
		panic(err)
	}
	db, err = DB.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(30)
	db.SetMaxOpenConns(100)

	if err := db.Ping(); err != nil {
		log.Printf("postgre connect error: %v", err)
		panic(err)
	}

	return func() {
		_ = db.Close()
	}
}

func connectionOptions() string {
	c := Conf.Postgre
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", c.Username, c.Password, c.Databases, c.Host, c.Port)
}
