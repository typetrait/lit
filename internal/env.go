package internal

type Environment struct {
	IsDebugEnabled bool
	DBHost         string
	DBUser         string
	DBPassword     string
	DBName         string
	DBPort         string
	S3Bucket       string
	LocalstackHost string
}
