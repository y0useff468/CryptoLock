package Connection

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=ep-plain-bar-a8mq0v2b-pooler.eastus2.azure.neon.tech user=neondb_owner password=npg_YZUfs5n8dcLm dbname=CryptoLock port=5432 sslmode=require"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	log.Println("Database connected successfully!")
}
