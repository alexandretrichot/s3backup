# s3backup

A simple CLI for backing up databases to an S3 bucket.

> **Important note**: At this time, only **MongoDB** is supported

## 🚀 Usage

### Backup MongoDB Database

To backup a MongoDB database with **docker**, use:

```shell
$ docker run -it --rm ghcr.io/alexandretrichot/s3backup:latest \
mongo backup -n "my-super-backup" \
--s3Region "<YOUR_REGION>" \
--s3AccessKeyId "<YOUR_ACCESS_KEY_ID>" \
--s3SecretKey "<YOUR_SECRET_KEY>" \
--s3Bucket "<YOUR_BUCKET>" \
--mongoUri "<YOUR_MONGO_URI>"
```
