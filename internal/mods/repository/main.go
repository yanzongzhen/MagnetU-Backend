package repository

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/yanzongzhen/magnetu/internal/mods/repository/api"
	"github.com/yanzongzhen/magnetu/internal/mods/repository/schema"
	"gorm.io/gorm"
)

type Repository struct {
	DB            *gorm.DB
	RepositoryAPI *api.Repository
	FolderAPI     *api.Folder
	FileAPI       *api.File
}

func (a *Repository) AutoMigrate(ctx context.Context) error {
	return a.DB.AutoMigrate(new(schema.Repository), new(schema.Folder), new(schema.File))
}

func (a *Repository) Init(ctx context.Context) error {
	if err := a.AutoMigrate(ctx); err != nil {
		return err
	}
	return nil
}

func (a *Repository) RegisterV1Routers(ctx context.Context, v1 *gin.RouterGroup) error {
	repository := v1.Group("repositories")
	{
		repository.GET("", a.RepositoryAPI.Query)
		repository.GET(":id", a.RepositoryAPI.Get)
		repository.POST("", a.RepositoryAPI.Create)
		repository.PUT(":id", a.RepositoryAPI.Update)
		repository.DELETE(":id", a.RepositoryAPI.Delete)
	}
	folder := v1.Group("folders")
	{
		folder.GET("", a.FolderAPI.Query)
		folder.GET(":id", a.FolderAPI.Get)
		folder.POST("", a.FolderAPI.Create)
		folder.PUT(":id", a.FolderAPI.Update)
		folder.DELETE(":id", a.FolderAPI.Delete)
	}
	file := v1.Group("files")
	{
		file.GET("", a.FileAPI.Query)
		file.GET(":id", a.FileAPI.Get)
		file.POST("", a.FileAPI.Create)
		file.PUT(":id", a.FileAPI.Update)
		file.DELETE(":id", a.FileAPI.Delete)
	}
	return nil
}

func (a *Repository) Release(ctx context.Context) error {
	return nil
}
