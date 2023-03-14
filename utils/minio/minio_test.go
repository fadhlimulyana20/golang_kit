package minio

import (
	"context"
	"testing"

	"github.com/minio/minio-go/v7"
)

func createClient() (*minio.Client, error) {
	endpoint := "localhost:9000"
	accessKey := "admin"
	accessSecret := "password"
	useSSL := false
	bucket := "golang-template"

	minio := NewMinioStorage(endpoint, accessKey, accessSecret, bucket, useSSL)
	client, err := minio.Client()
	return client, err
}

func TestClient(t *testing.T) {
	client, err := createClient()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(client)
}

func TestUpload(t *testing.T) {
	client, err := createClient()
	if err != nil {
		t.Fatal(err)
	}

	bucket := "golang-template"

	info, err := client.FPutObject(context.Background(), bucket, "text_test.txt", "./text_test.txt", minio.PutObjectOptions{ContentType: "application/txt"})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Successfully uploaded %s of size %d\n", "text_test.txt", info.Size)
}

// func createImage() *os.File {
// 	width := 200
// 	height := 100

// 	upLeft := image.Point{0, 0}
// 	lowRight := image.Point{width, height}

// 	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

// 	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
// 	cyan := color.RGBA{100, 200, 200, 0xff}

// 	// Set color for each pixel.
// 	for x := 0; x < width; x++ {
// 		for y := 0; y < height; y++ {
// 			switch {
// 			case x < width/2 && y < height/2: // upper left quadrant
// 				img.Set(x, y, cyan)
// 			case x >= width/2 && y >= height/2: // lower right quadrant
// 				img.Set(x, y, color.White)
// 			default:
// 				// Use zero value.
// 			}
// 		}
// 	}

// 	// Encode as PNG.
// 	f, _ := os.Create("./image.png")
// 	png.Encode(f, img)
// 	return f
// }

// func TestUploadMultipart(t *testing.T) {
// 	body := new(bytes.Buffer)
// 	writer := multipart.NewWriter(body)
// 	img := createImage()
// 	defer os.Remove(img.Name())

// 	part, err := writer.CreateFormFile("file", img.Name())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if _, err := io.Copy(part, img); err != nil {
// 		log.Fatal(err)
// 	}

// }
