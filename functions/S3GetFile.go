package functions

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func GetFile() {

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

	}
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}

	//getObject
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	objectCh := minioClient.ListObjects(ctx, "gamificacion", minio.ListObjectsOptions{
		Recursive: true,
	})
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
		fmt.Println(object)
	}

}
