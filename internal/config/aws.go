package config

import (
	"bytes"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

var awsSession *session.Session

func InitAWS() {
	conf := GetAWS()
	awsConfig := &aws.Config{
		Region:      aws.String(conf.DefaultRegion),
		Credentials: credentials.NewStaticCredentials(conf.AccessKeyID, conf.SecretAccessKey, ""),
	}
	var err error
	awsSession, err = session.NewSession(awsConfig)
	if err != nil {
		panic(err)
	}
}

func UploadFile(directory string, fileHeader *multipart.FileHeader) (string, string, error) {
	conf := GetAWS()
	fullpath := fileHeader.Filename
	fileNameWithoutExt := strings.TrimSuffix(fullpath, filepath.Ext(fullpath))
	slug := strings.ToLower(fileNameWithoutExt)
	slug = strings.ReplaceAll(slug, " ", "-")

	path := GetCurrentEnvironment() + "/" + directory + "/" + getCurrentDatePath() + "/" + slug + "." + getFileExtension(fullpath)

	file, err := fileHeader.Open()
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", "", err
	}

	input := &s3manager.UploadInput{
		Bucket:      aws.String(conf.Bucket),
		Key:         aws.String(path),
		Body:        bytes.NewReader(fileBytes),
		ContentType: aws.String(fileHeader.Header.Get("Content-Type")),
	}

	uploader := s3manager.NewUploader(awsSession)
	output, err := uploader.Upload(input)
	if err != nil {
		return "", "", err
	}

	return path, output.Location, nil
}

func UploadFileFromBytes(directory, filename string, fileBytes []byte) (string, string, error) {
	conf := GetAWS()
	extension := getFileExtension(filename)
	path := directory + "/" + uuid.New().String() + "." + extension
	contentType := ""

	switch extension {
	case "jpg", "jpeg":
		contentType = "image/jpeg"
	case "png":
		contentType = "image/png"
	default:
		return "", "", nil
	}

	input := &s3manager.UploadInput{
		Bucket:      aws.String(conf.Bucket),
		Key:         aws.String(path),
		Body:        bytes.NewReader(fileBytes),
		ContentType: aws.String(contentType),
	}

	uploader := s3manager.NewUploader(awsSession)
	output, err := uploader.Upload(input)
	if err != nil {
		return "", "", err
	}

	return path, output.Location, nil
}

func getFileExtension(str string) string {
	return str[strings.LastIndex(str, ".")+1:]
}

func getCurrentDatePath() string {
	now := time.Now()
	return now.Format("2006/1/2")
}

func GetFileTemporaryUrl(path string) (string, error) {
	conf := GetAWS()
	svc := s3.New(awsSession)
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(conf.Bucket),
		Key:    aws.String(path),
	})
	urlStr, err := req.Presign(time.Hour * 24)
	if err != nil {
		return "", err
	}
	return urlStr, nil
}
