package functions

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func RemoveFile(filename string) {
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

	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
		VersionID:        "myversionid",
	}
	err = minioClient.RemoveObject(context.Background(), "gamificacion", filename, opts)
	if err != nil {
		fmt.Println(err)
		return
	}
}
