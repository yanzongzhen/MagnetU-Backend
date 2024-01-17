package biz

import (
	"context"
	"github.com/yanzongzhen/magnetu/internal/mods/repository/dal"
	"github.com/yanzongzhen/magnetu/internal/mods/repository/schema"
	"github.com/yanzongzhen/magnetu/pkg/errors"
	"github.com/yanzongzhen/magnetu/pkg/util"
	"time"
)

// NetDisk management for NetDisk
type NetDisk struct {
	Trans         *util.Trans
	FileDAL       *dal.File
	RepositoryDAL *dal.Repository
}

// Query repositories from the data access object based on the provided parameters and options.
func (a *NetDisk) Query(ctx context.Context, params schema.FileQueryParam) (*schema.FileQueryResult, error) {
	params.Pagination = true

	result, err := a.FileDAL.Query(ctx, params, schema.FileQueryOptions{
		QueryOptions: util.QueryOptions{
			OrderFields: []util.OrderByParam{
				{Field: "created_at", Direction: util.DESC},
			},
		},
	})

	result.Data = result.Data.ToTree()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Get the specified NetDisk from the data access object.
func (a *NetDisk) Get(ctx context.Context, id string) (*schema.File, error) {
	file, err := a.FileDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if file == nil {
		return nil, errors.NotFound("", "File not found")
	}
	return file, nil
}

// Create a new File in the data access object.
func (a *NetDisk) Create(ctx context.Context, formItem *schema.FileForm) (*schema.File, error) {
	file := &schema.File{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(file); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.FileDAL.Create(ctx, file); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return file, nil
}

// Update the specified File in the data access object.
func (a *NetDisk) Update(ctx context.Context, id string, formItem *schema.FileForm) error {
	file, err := a.FileDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if file == nil {
		return errors.NotFound("", "File not found")
	}

	if err := formItem.FillTo(file); err != nil {
		return err
	}
	file.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.FileDAL.Update(ctx, file); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified File from the data access object.
func (a *NetDisk) Delete(ctx context.Context, id string) error {
	exists, err := a.FileDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "File not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.FileDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
