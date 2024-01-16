package repository

import (
	"github.com/google/wire"
	"github.com/yanzongzhen/magnetu/internal/mods/repository/api"
	"github.com/yanzongzhen/magnetu/internal/mods/repository/biz"
	"github.com/yanzongzhen/magnetu/internal/mods/repository/dal"
)

var Set = wire.NewSet(
	wire.Struct(new(Repository), "*"),
	wire.Struct(new(dal.Repository), "*"),
	wire.Struct(new(biz.Repository), "*"),
	wire.Struct(new(api.Repository), "*"),
	wire.Struct(new(dal.Folder), "*"),
	wire.Struct(new(dal.File), "*"),
)
