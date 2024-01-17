package mods

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/yanzongzhen/magnetu/internal/mods/cloud"
	"github.com/yanzongzhen/magnetu/internal/mods/rbac"
	"github.com/yanzongzhen/magnetu/internal/mods/repository"
	"github.com/yanzongzhen/magnetu/internal/mods/sys"
)

const (
	apiPrefix = "/api/"
)

// Collection of wire providers
var Set = wire.NewSet(
	wire.Struct(new(Mods), "*"),
	rbac.Set,
	sys.Set,
	repository.Set,
	cloud.Set,
)

type Mods struct {
	RBAC       *rbac.RBAC
	SYS        *sys.SYS
	Repository *repository.Repository
	CLOUD      *cloud.CLOUD
}

func (a *Mods) Init(ctx context.Context) error {
	if err := a.RBAC.Init(ctx); err != nil {
		return err
	}
	if err := a.SYS.Init(ctx); err != nil {
		return err
	}
	if err := a.Repository.Init(ctx); err != nil {
		return err
	}
	if err := a.CLOUD.Init(
		ctx,
	); err != nil {
		return err
	}

	return nil
}

func (a *Mods) RouterPrefixes() []string {
	return []string{
		apiPrefix,
	}
}

func (a *Mods) RegisterRouters(ctx context.Context, e *gin.Engine) error {
	gAPI := e.Group(apiPrefix)
	v1 := gAPI.Group("v1")

	if err := a.RBAC.RegisterV1Routers(ctx, v1); err != nil {
		return err
	}
	if err := a.SYS.RegisterV1Routers(ctx, v1); err != nil {
		return err
	}
	if err := a.Repository.RegisterV1Routers(ctx, v1); err != nil {
		return err
	}
	if err := a.CLOUD.RegisterV1Routers(ctx, v1); err != nil {
		return err
	}

	return nil
}

func (a *Mods) Release(ctx context.Context) error {
	if err := a.RBAC.Release(ctx); err != nil {
		return err
	}
	if err := a.SYS.Release(ctx); err != nil {
		return err
	}
	if err := a.Repository.
		Release(ctx); err != nil {
		return err
	}
	if err := a.CLOUD.Release(ctx); err != nil {
		return err
	}

	return nil
}
