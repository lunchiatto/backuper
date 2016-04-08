package stores

import (
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var bucketName = "lunchiatto-dev-backups"

// S3Store is a s3 store
type S3Store struct {
	svc *s3.S3
}

// CreateS3Store creates a S3 store
func CreateS3Store() *S3Store {
	// keyID := os.Getenv("AWS_ACCESS_KEY_ID")
	// key := os.Getenv("AWS_SECRET_ACCESS_KEY")
	svc := s3.New(session.New())

	return &S3Store{svc: svc}
}

// Upload uploads the file
func (s3Store *S3Store) Upload(r io.ReadSeeker) {
	uploadResult, err := s3Store.svc.PutObject(&s3.PutObjectInput{
		Body:   r,
		Bucket: &bucketName,
		Key:    fileName(),
	})
	fmt.Println(uploadResult)
	fmt.Println(err)
}

func fileName() *string {
	name := time.Now().Format("200601021504_dump.sql")
	return &name
}
