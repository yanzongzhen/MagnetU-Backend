package dal

import (
	"context"

	"github.com/yanzongzhen/magnetu/internal/mods/repository/schema"
	"github.com/yanzongzhen/magnetu/pkg/errors"
	"github.com/yanzongzhen/magnetu/pkg/util"
	"gorm.io/gorm"
)

// GetRepositoryDB Get repository storage instance
func GetRepositoryDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.Repository))
}

// Repository management for Repository
type Repository struct {
	DB *gorm.DB
}

// Query repositories from the database based on the provided parameters and options.
func (a *Repository) Query(ctx context.Context, params schema.RepositoryQueryParam, opts ...schema.RepositoryQueryOptions) (*schema.RepositoryQueryResult, error) {
	var opt schema.RepositoryQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetRepositoryDB(ctx, a.DB)
	if v := params.UserID; len(v) > 0 {
		db = db.Where("user_id = ?", v)
	}
	if v := params.CurrentCapacity; v != 0 {
		db = db.Where("current_capacity = ?", v)
	}
	if v := params.MaxCapacity; v != 0 {
		db = db.Where("max_capacity = ?", v)
	}
	if v := params.Permissions; len(v) > 0 {
		db = db.Where("permissions = ?", v)
	}

	var list schema.Repositories
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.RepositoryQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified repository from the database.
func (a *Repository) Get(ctx context.Context, id string, opts ...schema.RepositoryQueryOptions) (*schema.Repository, error) {
	var opt schema.RepositoryQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.Repository)
	ok, err := util.FindOne(ctx, GetRepositoryDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified repository exists in the database.
func (a *Repository) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetRepositoryDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new repository.
func (a *Repository) Create(ctx context.Context, item *schema.Repository) error {
	result := GetRepositoryDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified repository in the database.
func (a *Repository) Update(ctx context.Context, item *schema.Repository) error {
	result := GetRepositoryDB(ctx, a.DB).Where("id=?", item.RepositoryID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified repository from the database.
func (a *Repository) Delete(ctx context.Context, id string) error {
	result := GetRepositoryDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.Repository))
	return errors.WithStack(result.Error)
}
