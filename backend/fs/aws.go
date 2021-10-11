package fs

import (
	"bytes"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func getSession() *session.Session {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
	})
	return sess
}

func awsWrite(name string, content []byte) error {
	uploader := s3manager.NewUploader(getSession())
	filename := name + ".txt"
	file, _ := os.Create(filename)
	defer file.Close()
	file.WriteAt(content, int64(len(content)))
	file.Sync()
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key + "/" + filename),
		Body:   file,
	})
	return err
}

func awsRemove(name string) error {
	svc := s3.New(getSession())
	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key + "/" + name + ".txt"),
	})
	return err
}

func awsRead(name string) ([]byte, error) {
	filename := name + ".txt"
	file, _ := os.Create(filename)

	defer file.Close()

	downloader := s3manager.NewDownloader(getSession())
	_, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key + "/" + filename),
		})
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, file)
	data := buf.Bytes()
	return data, err
}

func awsGetAllFileNames() []string {
	svc := s3.New(getSession())
	resp, _ := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	})
	fileNames := []string{}
	for _, item := range resp.Contents {

		fileNames = append(fileNames, *item.Key)
	}
	return fileNames
}

func awsIsFileExist(name string) bool {
	_, err := awsRead(name)
	return err == nil
}
