package databases

import (
	"bytes"
	"os"
	"os/exec"
)

// PostgresBackup is a struct for backing up
type PostgresBackup struct {
	cmd    *exec.Cmd
	stdout *bytes.Buffer
	stderr *bytes.Buffer
}

// BuildPostgres creates an instance of PostgresBackup
func BuildPostgres() *PostgresBackup {
	cmd := buildCommand()
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	return &PostgresBackup{cmd: cmd, stderr: &stderr, stdout: &stdout}
}

// Run runs the backup
func (pb *PostgresBackup) Run() error {
	return pb.cmd.Run()
}

// Output returns backups result
func (pb *PostgresBackup) Output() string {
	return pb.stdout.String()
}

// Error returns possible errors
func (pb *PostgresBackup) Error() string {
	return pb.stderr.String()
}

func buildCommand() *exec.Cmd {
	containerName := os.Getenv("BACKUPER_CONTAINER")
	dBName := os.Getenv("BACKUPER_DB_NAME")
	return exec.Command("docker-compose", "run", "--rm", containerName, "pg_dump", "-U", "postgres", "-h", containerName, dBName)
}
