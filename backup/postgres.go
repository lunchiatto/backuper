package backup

import (
	"fmt"
	"strings"

	"github.com/lunchiatto/backuper/stores"
)

// Postgres store a PG backup
type Postgres struct {
	containerName string
}

// Run runs a backup
func Run() {
	store := stores.CreateS3Store()
	store.Upload(strings.NewReader("Hello World!"))
	fmt.Println(store)
}
