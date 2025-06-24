package database

import (
	"api-03/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb(){
	fmt.Println("Welcome")
	//dsn := "host=localhost port=5432 user=suntoh password=mypassword dbname=mydatabase sslmode=disable"

	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(postgres.Open(
        fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
            "suntoh", "mypassword", "localhost", "5432", "mydatabase", "disable"),
    ), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		panic("failed to connect database")
	}
	fmt.Println("Connected to the database successfully",db)
	db.AutoMigrate(&models.User{})
	
	Database = DbInstance{Db: db}
}