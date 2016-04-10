# Backuper

This package runs a backup and uploads it to S3.

It expects following env variables
* AWS_ACCESS_KEY_ID
* AWS_SECRET_ACCESS_KEY
* AWS_REGION
* BACKUPER_BUCKET - name of S3 Bucket
* BACKUPER_DB_NAME - Name of PG DB
* BACKUPER_CONTAINER - Name of docker container that runs the db
