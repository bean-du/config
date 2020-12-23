package conf

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	err   error
	db    *sql.DB
	DB    *gorm.DB
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

	return func() {
		_ = db.Close()
	}
}

func connectionOptions() string {
	c := Conf.Postgre
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", c.Username, c.Password, c.Databases, c.Host, c.Port)
}
