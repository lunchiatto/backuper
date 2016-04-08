package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/jasonlvhit/gocron"
	"github.com/lunchiatto/backuper/stores"
)

func main() {
	gocron.Every(1).Second().Do(run)

	<-gocron.Start()
}

func run() {
	// Abstract this out
	cmd := exec.Command("docker-compose", "run", "db", "pg_dump", "-U", "postgres", "-h", "db", "codequestmanager_development")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(stderr.String())
		fmt.Println(err)
		return
	}
	fmt.Printf("in all caps: %q\n", out.String())
	store := stores.CreateS3Store()
	// Make this return error
	store.Upload(strings.NewReader(out.String()))
}
