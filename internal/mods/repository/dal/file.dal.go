package dal

import (
	"context"

	"github.com/yanzongzhen/magnetu/internal/mods/repository/schema"
	"github.com/yanzongzhen/magnetu/pkg/errors"
	"github.com/yanzongzhen/magnetu/pkg/util"
	"gorm.io/gorm"
)

// GetFileDB Get file storage instance
func GetFileDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.File))
}

// File permissions for Repository
type File struct {
	DB *gorm.DB
}

// Query files from the database based on the provided parameters and options.
func (a *File) Query(ctx context.Context, params schema.FileQueryParam, opts ...schema.FileQueryOptions) (*schema.FileQueryResult, error) {
	var opt schema.FileQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetFileDB(ctx, a.DB)

	var list schema.Files
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.FileQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified file from the database.
func (a *File) Get(ctx context.Context, id string, opts ...schema.FileQueryOptions) (*schema.File, error) {
	var opt schema.FileQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.File)
	ok, err := util.FindOne(ctx, GetFileDB(ctx, a.DB).Where("file_id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified file exists in the database.
func (a *File) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetFileDB(ctx, a.DB).Where("file_id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new file.
func (a *File) Create(ctx context.Context, item *schema.File) error {
	result := GetFileDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified file in the database.
func (a *File) Update(ctx context.Context, item *schema.File) error {
	result := GetFileDB(ctx, a.DB).Where("file_id=?", item.FileID).Select("*").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified file from the database.
func (a *File) Delete(ctx context.Context, id string) error {
	result := GetFileDB(ctx, a.DB).Where("file_id=?", id).Delete(new(schema.File))
	return errors.WithStack(result.Error)
}

// DeleteByRepoID the specified file from the database.
func (a *File) DeleteByRepoID(ctx context.Context, id string) error {
	result := GetFileDB(ctx, a.DB).Where("repository_id=?", id).Delete(new(schema.File))
	return errors.WithStack(result.Error)
}
