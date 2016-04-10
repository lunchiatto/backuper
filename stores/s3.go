package stores

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3Store is a s3 store
type S3Store struct {
	svc *s3.S3
}

// CreateS3Store creates a S3 store
func CreateS3Store() *S3Store {
	return &S3Store{svc: s3.New(session.New())}
}

// Upload uploads the file
func (s3Store *S3Store) Upload(r io.ReadSeeker) error {
	bucketName := os.Getenv("BACKUPER_BUCKET")
	uploadResult, err := s3Store.svc.PutObject(&s3.PutObjectInput{
		Body:   r,
		Bucket: &bucketName,
		Key:    fileName(),
	})
	if err != nil {
		return err
	}
	fmt.Println("Successful upload")
	fmt.Println(uploadResult)
	return nil
}

func fileName() *string {
	name := time.Now().Format("200601021504_dump.sql")
	return &name
}
