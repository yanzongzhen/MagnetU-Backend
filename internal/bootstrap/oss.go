package bootstrap

import (
	"github.com/yanzongzhen/magnetu/internal/config"
	"github.com/yanzongzhen/magnetu/pkg/oss"
)

func InitOSS() error {
	if config.C.EnableOSS() {
		if config.C.OSS.Minio.Domain != "" {
			minioC, err := oss.NewMinioClient(oss.MinioClientConfig{
				Domain:          config.C.OSS.Minio.Domain,
				Endpoint:        config.C.OSS.Minio.Endpoint,
				AccessKeyID:     config.C.OSS.Minio.AccessKeyID,
				SecretAccessKey: config.C.OSS.Minio.SecretAccessKey,
				BucketName:      config.C.OSS.BucketName,
				Prefix:          config.C.OSS.Prefix,
			})
			if err != nil {
				return err
			}
			oss.SetGlobal(func() oss.IClient {
				return minioC
			})
		}
		if config.C.OSS.S3.Domain != "" {
			s3C, err := oss.NewS3Client(oss.S3ClientConfig{
				Domain:          config.C.OSS.S3.Domain,
				Region:          config.C.OSS.S3.Region,
				AccessKeyID:     config.C.OSS.S3.AccessKeyID,
				SecretAccessKey: config.C.OSS.S3.SecretAccessKey,
				BucketName:      config.C.OSS.BucketName,
				Prefix:          config.C.OSS.Prefix,
			})
			if err != nil {
				return err
			}
			oss.SetGlobal(func() oss.IClient {
				return s3C
			})
		}
	}
	return nil
}
