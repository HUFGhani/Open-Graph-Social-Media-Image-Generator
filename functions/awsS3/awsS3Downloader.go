package awsS3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
)

func S3downloadAssets() {
	log.Printf("aws congfig")

	bucketname := os.Getenv("S3_BUCKET_NAME")
	files := []string{"background.jpeg", "DrSugiyama-Regular.ttf", "OpenSans-Bold.ttf"}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-2")},
	)

	if err != nil {
		log.Printf("unable to connect %v", err)
	}

	downloader := s3manager.NewDownloader(sess)

	for _, file := range files {
		downloadFiles(bucketname, file, downloader)
	}
}

func downloadFiles(bucketname string, file string, downloader *s3manager.Downloader) {
	tmpFilePath, err := os.Create("/tmp/" + file)

	if err != nil {
		log.Printf("Unable to open file %q, %v", file, err)
	}

	defer tmpFilePath.Close()

	_, err = downloader.Download(tmpFilePath, &s3.GetObjectInput{
		Bucket: aws.String(bucketname),
		Key:    aws.String(file),
	})

	if err != nil {
		log.Printf("Unable to download s3File %q, %v", file, err)
	}
}
