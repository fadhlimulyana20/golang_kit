package minio

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
)

type minioStorage struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
	BucketName      string
}

type MinioStorageContract interface {
	Client() (*minio.Client, error)
	// Upload multipart file using go routine
	UploadMultipart(fileUploadedPath chan string, e chan error, fileHeader *multipart.FileHeader, pathFile string)
	// Generate temporary public URL. It will expire within 24 hours.
	GetTemporaryPublicUrl(filePath string) (*url.URL, error)
	// Delete file from bucket
	DeleteFile(filepath string) error
}

func NewMinioStorage(endpoint, accessKeyID, secretAccessKey, bucket string, useSSL bool) MinioStorageContract {
	return &minioStorage{
		Endpoint:        endpoint,
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
		UseSSL:          useSSL,
		BucketName:      bucket,
	}
}

func (m *minioStorage) Client() (*minio.Client, error) {
	minioClient, err := minio.New(m.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(m.AccessKeyID, m.SecretAccessKey, ""),
	})

	if err != nil {
		logrus.Error(err)
		return minioClient, err
	}

	logrus.Infof("%#v\n", minioClient)
	return minioClient, nil
}

func (m *minioStorage) generateMultipartTempFilePath(fileHeader *multipart.FileHeader) (string, error) {
	multipartFile, err := fileHeader.Open()
	if err != nil {
		return "", err
	}

	file, err := os.CreateTemp("./", fmt.Sprintf("file_*%s", filepath.Ext(fileHeader.Filename)))
	if err != nil {
		return "", err
	}

	_, err = io.Copy(file, multipartFile)
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

func (m *minioStorage) UploadMultipart(fileUploadedPath chan string, e chan error, fileHeader *multipart.FileHeader, pathFile string) {
	file, err := m.generateMultipartTempFilePath(fileHeader)
	defer os.Remove(file)
	if err != nil {
		fileUploadedPath <- ""
		e <- err
		return
	}

	client, err := m.Client()
	if err != nil {
		e <- err
		return
	}

	filename := fmt.Sprintf("file_%d%s", time.Now().Unix(), filepath.Ext(fileHeader.Filename))
	destPath := fmt.Sprintf("%s/%s", pathFile, filename)

	_, err = client.FPutObject(context.Background(), m.BucketName, destPath, file, minio.PutObjectOptions{})
	if err != nil {
		fileUploadedPath <- ""
		e <- err
		return
	}
	e <- nil
	fileUploadedPath <- destPath
}

func (m *minioStorage) GetTemporaryPublicUrl(filePath string) (*url.URL, error) {
	filePathElem := strings.Split(filePath, "/")
	filename := filePathElem[len(filePathElem)-1]
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	client, err := m.Client()
	if err != nil {
		return nil, nil
	}

	presignedURL, err := client.PresignedGetObject(context.Background(), m.BucketName, filePath, time.Second*24*60*60, reqParams)
	if err != nil {
		return nil, nil
	}

	return presignedURL, nil
}

func (m *minioStorage) DeleteFile(filepath string) error {
	client, err := m.Client()
	if err != nil {
		return err
	}

	err = client.RemoveObject(context.Background(), m.BucketName, filepath, minio.RemoveObjectOptions{})
	return err
}
