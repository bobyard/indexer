package main

import (
	"github.com/bobyard/indexer/db"
	"github.com/bobyard/indexer/models"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	bobYardDB := os.Getenv("BOBYARD")
	db.Connect(bobYardDB)

	err = db.Engine.Sync2(new(models.Activities), new(models.Offers), new(models.Collections), new(models.Tokens), new(models.Lists), new(models.Domains), new(models.Orders))
	if err != nil {
		panic(err)
	}
}
