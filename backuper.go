package main

import (
	"fmt"
	"strings"

	"github.com/jasonlvhit/gocron"
	"github.com/lunchiatto/backuper/databases"
	"github.com/lunchiatto/backuper/stores"
)

func main() {
	gocron.Every(1).Second().Do(run)

	<-gocron.Start()
}

func run() {
	postgres := databases.BuildPostgres()
	err := postgres.Run()
	if err != nil {
		fmt.Println("Sth went wrong 🙊")
		fmt.Println(err)
		fmt.Println(postgres.Error())
		return
	}
	store := stores.CreateS3Store()
	// Make this return error
	if err := store.Upload(strings.NewReader(postgres.Output())); err != nil {
		fmt.Println("AWS upload error")
		fmt.Println(err)
	}
}
