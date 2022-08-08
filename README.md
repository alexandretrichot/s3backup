# s3backup

A simple CLI for backing up databases to an S3 bucket.

> **Important note**: At this time, only **MongoDB** is supported

## 🚀 Usage

```shell
$ s3backup mongo backup -h
MongoDB Backup

Usage:
  s3backup mongo backup [flags]

Flags:
  -h, --help   help for backup

Global Flags:
  -u, --mongoUri string        The Database URI string (env: MONGO_URI)
      --s3AccessKeyId string   The s3 access key id (env: S3_ACCESS_KEY_ID)
      --s3Bucket string        The name of the bucket (env: S3_BUCKET)
      --s3Endpoint string      The s3 endpoint URL (env: S3_ENDPOINT)
      --s3Region string        The region to use for the backup (env: S3_REGION)
      --s3SecretKey string     The s3 secret key (env: S3_SECRET_KEY)
```
