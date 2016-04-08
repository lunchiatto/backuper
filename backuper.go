package main

import (
	"github.com/jasonlvhit/gocron"
	"github.com/lunchiatto/backuper/backup"
)

func main() {
	gocron.Every(1).Second().Do(backup.Run)

	<-gocron.Start()
}
