package biz

import (
	"context"
	"github.com/yanzongzhen/magnetu/internal/mods/repository/dal"
	"github.com/yanzongzhen/magnetu/internal/mods/repository/schema"
	"time"

	"github.com/yanzongzhen/magnetu/pkg/errors"
	"github.com/yanzongzhen/magnetu/pkg/util"
)

// Folder management for Folder
type Folder struct {
	Trans     *util.Trans
	FolderDAL *dal.Folder
}

// Query repositories from the data access object based on the provided parameters and options.
func (a *Folder) Query(ctx context.Context, params schema.FolderQueryParam) (*schema.FolderQueryResult, error) {
	params.Pagination = true

	result, err := a.FolderDAL.Query(ctx, params, schema.FolderQueryOptions{
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

// Get the specified Folder from the data access object.
func (a *Folder) Get(ctx context.Context, id string) (*schema.Folder, error) {
	Folder, err := a.FolderDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if Folder == nil {
		return nil, errors.NotFound("", "Folder not found")
	}
	return Folder, nil
}

// Create a new Folder in the data access object.
func (a *Folder) Create(ctx context.Context, formItem *schema.FolderForm) (*schema.Folder, error) {
	Folder := &schema.Folder{
		FolderID:  util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(Folder); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.FolderDAL.Create(ctx, Folder); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return Folder, nil
}

// Update the specified Folder in the data access object.
func (a *Folder) Update(ctx context.Context, id string, formItem *schema.FolderForm) error {
	Folder, err := a.FolderDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if Folder == nil {
		return errors.NotFound("", "Folder not found")
	}

	if err := formItem.FillTo(Folder); err != nil {
		return err
	}
	Folder.CreatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.FolderDAL.Update(ctx, Folder); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified Folder from the data access object.
func (a *Folder) Delete(ctx context.Context, id string) error {
	exists, err := a.FolderDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Folder not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.FolderDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
