package biz

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/yanzongzhen/magnetu/internal/config"
	"github.com/yanzongzhen/magnetu/pkg/encoding/json"
	"github.com/yanzongzhen/magnetu/pkg/oss"
	"log"
	"time"

	"github.com/yanzongzhen/magnetu/internal/mods/cloud/dal"
	"github.com/yanzongzhen/magnetu/internal/mods/cloud/schema"
	"github.com/yanzongzhen/magnetu/pkg/errors"
	"github.com/yanzongzhen/magnetu/pkg/util"
)

// Cloud file management
type Cloud struct {
	Trans    *util.Trans
	CloudDAL *dal.Cloud
}

// Upload file to server
func (a *Cloud) Upload(ctx *gin.Context) ([]*schema.Cloud, error) {
	fileRes := make([]*schema.Cloud, 0)
	if !config.C.EnableOSS() {
		return nil, errors.NotFound("", "OSS is not enabled")
	}
	form, _ := ctx.MultipartForm()
	files := form.File["file"]
	contentType := "application/octet-stream"
	for _, file := range files {
		fileOpen, _ := file.Open()
		info, err := oss.Ins.PutObject(ctx, config.C.OSS.BucketName, config.C.OSS.Prefix+"/"+file.Filename,
			fileOpen, file.Size, oss.PutObjectOptions{ContentType: contentType})
		if err != nil {
			log.Printf("Failed to upload %s of size %d \t\t %s", config.C.OSS.Prefix+"/"+file.Filename, file.Size, err.Error())
		}
		//info.Size 文件大小，以为B为单位
		log.Printf("Successfully uploaded %s of size : \t\t %dB \t %.2fKB\t %.2fMB\t", config.C.OSS.Prefix+"/"+file.Filename, info.Size, float64(info.Size/1024.0), float64(info.Size/1024.0/1024.0))
		infoMeta, _ := json.Marshal(info)
		clu := schema.Cloud{
			ID:        util.NewXID(),
			CreatedAt: time.Now(),
			Name:      file.Filename,
			Size:      info.Size,
			URL:       info.URL,
			OSS:       info.URL,
			FileMeta:  string(infoMeta),
		}
		err = a.CloudDAL.Create(ctx, &clu)
		if err != nil {
			log.Printf("Failed to save %s of size %d \t\t %s", config.C.OSS.Prefix+"/"+file.Filename, file.Size, err.Error())
		}
		fileRes = append(fileRes, &clu)
	}
	return fileRes, nil
}

// Query clouds from the data access object based on the provided parameters and options.
func (a *Cloud) Query(ctx context.Context, params schema.CloudQueryParam) (*schema.CloudQueryResult, error) {
	params.Pagination = true

	result, err := a.CloudDAL.Query(ctx, params, schema.CloudQueryOptions{
		QueryOptions: util.QueryOptions{
			OrderFields: []util.OrderByParam{
				{Field: "created_at", Direction: util.DESC},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Get the specified cloud from the data access object.
func (a *Cloud) Get(ctx context.Context, id string) (*schema.Cloud, error) {
	cloud, err := a.CloudDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if cloud == nil {
		return nil, errors.NotFound("", "Cloud not found")
	}
	return cloud, nil
}

// Create a new cloud in the data access object.
func (a *Cloud) Create(ctx context.Context, formItem *schema.CloudForm) (*schema.Cloud, error) {
	cloud := &schema.Cloud{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(cloud); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.CloudDAL.Create(ctx, cloud); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return cloud, nil
}

// Update the specified cloud in the data access object.
func (a *Cloud) Update(ctx context.Context, id string, formItem *schema.CloudForm) error {
	cloud, err := a.CloudDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if cloud == nil {
		return errors.NotFound("", "Cloud not found")
	}

	if err := formItem.FillTo(cloud); err != nil {
		return err
	}
	cloud.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.CloudDAL.Update(ctx, cloud); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified cloud from the data access object.
func (a *Cloud) Delete(ctx context.Context, id string) error {
	exists, err := a.CloudDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Cloud not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.CloudDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
