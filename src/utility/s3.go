// utility/s3.go
package utility

import (
	"context"
	"fmt"
	"io"

	awsCfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/devopscorner/golang-restfulapi-bedrock/src/config"
)

func InitS3Client(cfg *config.Config) (*s3.Client, error) {
	s3Cfg, err := awsCfg.LoadDefaultConfig(context.TODO(),
		awsCfg.WithRegion(cfg.AWSRegion),
		awsCfg.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.AWSAccessKey,
			cfg.AWSSecretKey,
			"",
		)),
	)
	if err != nil {
		return nil, err
	}

	return s3.NewFromConfig(s3Cfg), nil
}

// UploadFileToS3 uploads a file to S3 and returns the URL
func UploadFileToS3(ctx context.Context, client *s3.Client, bucket string, key string, file io.Reader) (string, error) {
	_, err := client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   file,
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, key), nil
}
