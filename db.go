package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	//dsn := "host=localhost user=postgres password=1234 dbname=league port=5432 sslmode=disable"
	dsn := "postgresql://postgres:cQIxBPSxGSJnGCWlbxCYamXaqEzJIwDV@shuttle.proxy.rlwy.net:19637/railway"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Database connection failed: " + err.Error())
	}
	err = DB.AutoMigrate(&Team{}, &Match{})
	if err != nil {
		panic("Tables could not be created: " + err.Error())
	}
	fmt.Println("Database connection and tables are ready!")
}
