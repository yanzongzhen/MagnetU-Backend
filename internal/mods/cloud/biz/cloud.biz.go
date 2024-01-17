package biz

import (
	"context"
	"time"

	"github.com/yanzongzhen/magnetu/internal/mods/cloud/dal"
	"github.com/yanzongzhen/magnetu/internal/mods/cloud/schema"
	"github.com/yanzongzhen/magnetu/pkg/errors"
	"github.com/yanzongzhen/magnetu/pkg/util"
)

// Cloud file management
type Cloud struct {
	Trans    *util.Trans
	CloudDAL *dal.Cloud
}

// Query clouds from the data access object based on the provided parameters and options.
func (a *Cloud) Query(ctx context.Context, params schema.CloudQueryParam) (*schema.CloudQueryResult, error) {
	params.Pagination = true

	result, err := a.CloudDAL.Query(ctx, params, schema.CloudQueryOptions{
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

// Get the specified cloud from the data access object.
func (a *Cloud) Get(ctx context.Context, id string) (*schema.Cloud, error) {
	cloud, err := a.CloudDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if cloud == nil {
		return nil, errors.NotFound("", "Cloud not found")
	}
	return cloud, nil
}

// Create a new cloud in the data access object.
func (a *Cloud) Create(ctx context.Context, formItem *schema.CloudForm) (*schema.Cloud, error) {
	cloud := &schema.Cloud{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(cloud); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.CloudDAL.Create(ctx, cloud); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return cloud, nil
}

// Update the specified cloud in the data access object.
func (a *Cloud) Update(ctx context.Context, id string, formItem *schema.CloudForm) error {
	cloud, err := a.CloudDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if cloud == nil {
		return errors.NotFound("", "Cloud not found")
	}

	if err := formItem.FillTo(cloud); err != nil {
		return err
	}
	cloud.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.CloudDAL.Update(ctx, cloud); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified cloud from the data access object.
func (a *Cloud) Delete(ctx context.Context, id string) error {
	exists, err := a.CloudDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Cloud not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.CloudDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
