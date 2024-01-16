package biz

import (
	"context"
	"time"

	"github.com/yanzongzhen/magnetu/internal/mods/repository/dal"
	"github.com/yanzongzhen/magnetu/internal/mods/repository/schema"
	"github.com/yanzongzhen/magnetu/pkg/errors"
	"github.com/yanzongzhen/magnetu/pkg/util"
)

// Repository management for Repository
type Repository struct {
	Trans         *util.Trans
	RepositoryDAL *dal.Repository
}

// Query repositories from the data access object based on the provided parameters and options.
func (a *Repository) Query(ctx context.Context, params schema.RepositoryQueryParam) (*schema.RepositoryQueryResult, error) {
	params.Pagination = true

	result, err := a.RepositoryDAL.Query(ctx, params, schema.RepositoryQueryOptions{
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

// Get the specified repository from the data access object.
func (a *Repository) Get(ctx context.Context, id string) (*schema.Repository, error) {
	repository, err := a.RepositoryDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if repository == nil {
		return nil, errors.NotFound("", "Repository not found")
	}
	return repository, nil
}

// Create a new repository in the data access object.
func (a *Repository) Create(ctx context.Context, formItem *schema.RepositoryForm) (*schema.Repository, error) {
	repository := &schema.Repository{
		RepositoryID: util.NewXID(),
		CreatedAt:    time.Now(),
	}

	if err := formItem.FillTo(repository); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.RepositoryDAL.Create(ctx, repository); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return repository, nil
}

// Update the specified repository in the data access object.
func (a *Repository) Update(ctx context.Context, id string, formItem *schema.RepositoryForm) error {
	repository, err := a.RepositoryDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if repository == nil {
		return errors.NotFound("", "Repository not found")
	}

	if err := formItem.FillTo(repository); err != nil {
		return err
	}
	repository.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.RepositoryDAL.Update(ctx, repository); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified repository from the data access object.
func (a *Repository) Delete(ctx context.Context, id string) error {
	exists, err := a.RepositoryDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Repository not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.RepositoryDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
