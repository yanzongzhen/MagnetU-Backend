package dal

import (
	"context"

	"github.com/yanzongzhen/magnetu/internal/mods/repository/schema"
	"github.com/yanzongzhen/magnetu/pkg/errors"
	"github.com/yanzongzhen/magnetu/pkg/util"
	"gorm.io/gorm"
)

// GetFolderDB Get folder storage instance
func GetFolderDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.Folder))
}

// Folder permissions for Repository
type Folder struct {
	DB *gorm.DB
}

// Query folders from the database based on the provided parameters and options.
func (a *Folder) Query(ctx context.Context, params schema.FolderQueryParam, opts ...schema.FolderQueryOptions) (*schema.FolderQueryResult, error) {
	var opt schema.FolderQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetFolderDB(ctx, a.DB)
	if v := params.FolderName; len(v) > 0 {
		db = db.Where("folder_name = ?", v)
	}

	var list schema.Folders
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.FolderQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified folder from the database.
func (a *Folder) Get(ctx context.Context, id string, opts ...schema.FolderQueryOptions) (*schema.Folder, error) {
	var opt schema.FolderQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.Folder)
	ok, err := util.FindOne(ctx, GetFolderDB(ctx, a.DB).Where("folder_id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified folder exists in the database.
func (a *Folder) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetFolderDB(ctx, a.DB).Where("folder_id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new folder.
func (a *Folder) Create(ctx context.Context, item *schema.Folder) error {
	result := GetFolderDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified folder in the database.
func (a *Folder) Update(ctx context.Context, item *schema.Folder) error {
	result := GetFolderDB(ctx, a.DB).Where("folder_id=?", item.FolderID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified folder from the database.
func (a *Folder) Delete(ctx context.Context, id string) error {
	result := GetFolderDB(ctx, a.DB).Where("folder_id=?", id).Delete(new(schema.Folder))
	return errors.WithStack(result.Error)
}
