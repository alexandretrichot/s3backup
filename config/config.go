package config

type Config struct {
	MongoURI      string `env:"MONGO_URI"`
	S3Endpoint    string `env:"S3_ENDPOINT"`
	S3AccessKeyId string `env:"S3_ACCESS_KEY_ID"`
	S3SecretKey   string `env:"S3_SECRET_KEY"`
	S3Bucket      string `env:"S3_BUCKET"`
	S3Region      string `env:"S3_REGION"`

	MongodumpPath string `env:"MONGODUMP_PATH" envDefault:"mongodump"`
}
