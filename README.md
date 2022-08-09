# s3backup

A simple CLI for backing up databases to an S3 bucket.

> **Important note**: At this time, only **MongoDB** is supported

## 🚀 Usage

### Backup MongoDB Database

To manual backup a MongoDB database with **docker**, use:

```shell
$ docker run -it --rm ghcr.io/alexandretrichot/s3backup:latest \
mongo backup -n "my-super-backup" \
--s3Region "<YOUR_REGION>" \
--s3AccessKeyId "<YOUR_ACCESS_KEY_ID>" \
--s3SecretKey "<YOUR_SECRET_KEY>" \
--s3Bucket "<YOUR_BUCKET>" \
--mongoUri "<YOUR_MONGO_URI>"
```

To manual backup a MongoDB database on **k8s**, use:

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: manual-backup
spec:
  ttlSecondsAfterFinished: 300
  template:
    spec:
      containers:
        - name: backup
          image: ghcr.io/alexandretrichot/s3backup:latest
          args:
            - 'mongo'
            - 'backup'
            - '-n'
            - 'myapp-manual-backup'
          env:
            - name: MONGO_URI
              value: '<YOUR_MONGO_URI>'
            - name: S3_REGION
              value: '<YOUR_REGION>'
            - name: S3_ACCESS_KEY_ID
              value: '<YOUR_ACCESS_KEY_ID>'
            - name: S3_SECRET_KEY
              value: '<YOUR_SECRET_KEY>'
            - name: S3_BUCKET
              value: '<YOUR_BUCKET>'
      restartPolicy: Never
```

Or with a **CronJob**

```yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: auto-backup
spec:
  schedule: '0 2 * * *'
  jobTemplate:
    spec:
      ttlSecondsAfterFinished: 300
      template:
        spec:
          containers:
            - name: backup
              image: ghcr.io/alexandretrichot/s3backup:latest
              args:
                - 'mongo'
                - 'backup'
                - '-n'
                - 'myapp-auto-backup'
              env:
                - name: MONGO_URI
                  value: '<YOUR_MONGO_URI>'
                - name: S3_REGION
                  value: '<YOUR_REGION>'
                - name: S3_ACCESS_KEY_ID
                  value: '<YOUR_ACCESS_KEY_ID>'
                - name: S3_SECRET_KEY
                  value: '<YOUR_SECRET_KEY>'
                - name: S3_BUCKET
                  value: '<YOUR_BUCKET>'
          restartPolicy: OnFailure
```
