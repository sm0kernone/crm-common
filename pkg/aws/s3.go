package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func Download(awsSession *session.Session, bucket, filePath string) ([]byte, error) {
	downloader := s3manager.NewDownloader(awsSession)

	awsBuffer := aws.WriteAtBuffer{}
	_, err := downloader.Download(&awsBuffer, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filePath),
	})

	return awsBuffer.Bytes(), err
}

func GetAwsSession(region, accessKeyID, secretKey string) *session.Session {
	sess, err := session.NewSession(
		&aws.Config{
			Region:      aws.String(region),
			Credentials: credentials.NewStaticCredentials(accessKeyID, secretKey, ""),
		},
	)
	if err != nil {
		panic(err)
	}
	return sess
}
