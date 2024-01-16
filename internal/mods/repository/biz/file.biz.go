package biz

import (
	"context"
	"github.com/yanzongzhen/magnetu/internal/mods/repository/dal"
	"github.com/yanzongzhen/magnetu/internal/mods/repository/schema"
	"time"

	"github.com/yanzongzhen/magnetu/pkg/errors"
	"github.com/yanzongzhen/magnetu/pkg/util"
)

// File management for Folder
type File struct {
	Trans   *util.Trans
	FileDAL *dal.File
}

// Query repositories from the data access object based on the provided parameters and options.
func (a *File) Query(ctx context.Context, params schema.FileQueryParam) (*schema.FileQueryResult, error) {
	params.Pagination = true

	result, err := a.FileDAL.Query(ctx, params, schema.FileQueryOptions{
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

// Get the specified File from the data access object.
func (a *File) Get(ctx context.Context, id string) (*schema.File, error) {
	File, err := a.FileDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if File == nil {
		return nil, errors.NotFound("", "File not found")
	}
	return File, nil
}

// Create a new File in the data access object.
func (a *File) Create(ctx context.Context, formItem *schema.FileForm) (*schema.File, error) {
	File := &schema.File{
		FileID:     util.NewXID(),
		UploadTime: time.Now(),
	}

	if err := formItem.FillTo(File); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.FileDAL.Create(ctx, File); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return File, nil
}

// Update the specified File in the data access object.
func (a *File) Update(ctx context.Context, id string, formItem *schema.FileForm) error {
	File, err := a.FileDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if File == nil {
		return errors.NotFound("", "File not found")
	}

	if err := formItem.FillTo(File); err != nil {
		return err
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.FileDAL.Update(ctx, File); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified File from the data access object.
func (a *File) Delete(ctx context.Context, id string) error {
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
