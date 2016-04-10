package main

import (
	"fmt"
	"os"

	"github.com/jasonlvhit/gocron"
	"github.com/kruszczynski/barkup"
)

func main() {
	gocron.Every(5).Minutes().Do(run)

	<-gocron.Start()
}

func run() {
	postgres := &barkup.Postgres{
		Host:     os.Getenv("BACKUPER_CONTAINER"),
		Port:     os.Getenv("BACKUPER_DB_PORT"),
		DB:       os.Getenv("BACKUPER_DB_NAME"),
		Username: os.Getenv("BACKUPER_DB_USER"),
	}

	// Configure a S3 storer
	s3 := &barkup.S3{
		Region:       os.Getenv("AWS_REGION"),
		Bucket:       os.Getenv("BACKUPER_BUCKET"),
		AccessKey:    os.Getenv("AWS_ACCESS_KEY_ID"),
		ClientSecret: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}

	if err := postgres.Export().To("/", s3); err != nil {
		fmt.Println(err)
	}
}
