package minio

import (
	"context"
	"testing"

	"github.com/minio/minio-go/v7"
)

func createClient() MinioStorageContract {
	endpoint := "localhost:9000"
	accessKey := "admin"
	accessSecret := "password"
	useSSL := false
	bucket := "golang-template"

	m := NewMinioStorage(endpoint, accessKey, accessSecret, bucket, useSSL)
	return m
}

var filename = "text_test.txt"

func TestClient(t *testing.T) {
	m := createClient()
	client, err := m.Client()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(client)
}

func TestUpload(t *testing.T) {
	m := createClient()
	client, err := m.Client()
	if err != nil {
		t.Fatal(err)
	}

	bucket := "golang-template"

	info, err := client.FPutObject(context.Background(), bucket, filename, "./text_test.txt", minio.PutObjectOptions{ContentType: "application/txt"})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Successfully uploaded %s of size %d\n", filename, info.Size)
}

func TestGetPresignedURL(t *testing.T) {
	m := createClient()

	url, err := m.GetTemporaryPublicUrl(filename)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(url)
}

func TestDeleteFile(t *testing.T) {
	m := createClient()

	if err := m.DeleteFile(filename); err != nil {
		t.Fatal(err)
	}

	t.Log("File deleted")
}
