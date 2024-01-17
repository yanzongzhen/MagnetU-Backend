package cloud

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/yanzongzhen/magnetu/internal/mods/cloud/api"
	"github.com/yanzongzhen/magnetu/internal/mods/cloud/schema"
	"gorm.io/gorm"
)

type CLOUD struct {
	DB       *gorm.DB
	CloudAPI *api.Cloud
}

func (a *CLOUD) AutoMigrate(ctx context.Context) error {
	return a.DB.AutoMigrate(new(schema.Cloud))
}

func (a *CLOUD) Init(ctx context.Context) error {
	if err := a.AutoMigrate(ctx); err != nil {
		return err
	}
	return nil
}

func (a *CLOUD) RegisterV1Routers(ctx context.Context, v1 *gin.RouterGroup) error {
	cloud := v1.Group("clouds")
	{
		cloud.POST("upload", a.CloudAPI.Upload)
		cloud.GET("", a.CloudAPI.Query)
		cloud.GET(":id", a.CloudAPI.Get)
		cloud.POST("", a.CloudAPI.Create)
		cloud.PUT(":id", a.CloudAPI.Update)
		cloud.DELETE(":id", a.CloudAPI.Delete)
	}
	return nil
}

func (a *CLOUD) Release(ctx context.Context) error {
	return nil
}
