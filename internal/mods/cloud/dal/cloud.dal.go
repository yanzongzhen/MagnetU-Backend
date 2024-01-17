package dal

import (
	"context"

	"github.com/yanzongzhen/magnetu/internal/mods/cloud/schema"
	"github.com/yanzongzhen/magnetu/pkg/errors"
	"github.com/yanzongzhen/magnetu/pkg/util"
	"gorm.io/gorm"
)

// GetCloudDB Get cloud storage instance
func GetCloudDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.Cloud))
}

// Cloud file management
type Cloud struct {
	DB *gorm.DB
}

// Query clouds from the database based on the provided parameters and options.
func (a *Cloud) Query(ctx context.Context, params schema.CloudQueryParam, opts ...schema.CloudQueryOptions) (*schema.CloudQueryResult, error) {
	var opt schema.CloudQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetCloudDB(ctx, a.DB)

	var list schema.Clouds
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.CloudQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified cloud from the database.
func (a *Cloud) Get(ctx context.Context, id string, opts ...schema.CloudQueryOptions) (*schema.Cloud, error) {
	var opt schema.CloudQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.Cloud)
	ok, err := util.FindOne(ctx, GetCloudDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified cloud exists in the database.
func (a *Cloud) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetCloudDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new cloud.
func (a *Cloud) Create(ctx context.Context, item *schema.Cloud) error {
	result := GetCloudDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified cloud in the database.
func (a *Cloud) Update(ctx context.Context, item *schema.Cloud) error {
	result := GetCloudDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified cloud from the database.
func (a *Cloud) Delete(ctx context.Context, id string) error {
	result := GetCloudDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.Cloud))
	return errors.WithStack(result.Error)
}
