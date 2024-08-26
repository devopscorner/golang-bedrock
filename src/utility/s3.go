// utility/s3.go
package utility

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/devopscorner/golang-bedrock/src/config"
)

func InitS3Client(cfg *config.Config) (*s3.Client, error) {
	// Create a custom AWS config with static credentials
	awsCfg := aws.Config{
		Region: cfg.AWSRegion,
		Credentials: credentials.NewStaticCredentialsProvider(
			cfg.AWSAccessKey,
			cfg.AWSSecretKey,
			"",
		),
	}

	// Create and return the S3 client
	return s3.NewFromConfig(awsCfg), nil
}

// UploadFileToS3 uploads a file to S3 and returns the URL
func UploadFileToS3(ctx context.Context, client *s3.Client, bucket string, key string, file io.Reader) (string, error) {
	_, err := client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, key), nil
}
