package functions

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func CreateFile(filename string) {
	/* ctx := context.Background() */
	endpoint := "sfo2.digitaloceanspaces.com"
	accessKeyID := "2IC4XIQLNKPBZB5IDWIL"
	secretAccessKey := "bAf2qiJeQ8LzF0GjsOwG1b1U01GVwI5wlFYSZD0rCDA"
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Println(err)
	}
	buckets, err := minioClient.ListBuckets(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}

	//UPLOAD FILE
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	uploadInfo, err := minioClient.PutObject(context.Background(), "gamificacion", "generateGif", file, fileStat.Size(), minio.PutObjectOptions{ContentType: "image/gif"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully uploaded bytes: ", uploadInfo)

}
