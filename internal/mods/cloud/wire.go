package cloud

import (
	"github.com/google/wire"
	"github.com/yanzongzhen/magnetu/internal/mods/cloud/api"
	"github.com/yanzongzhen/magnetu/internal/mods/cloud/biz"
	"github.com/yanzongzhen/magnetu/internal/mods/cloud/dal"
)

var Set = wire.NewSet(
	wire.Struct(new(CLOUD), "*"),
	wire.Struct(new(dal.Cloud), "*"),
	wire.Struct(new(biz.Cloud), "*"),
	wire.Struct(new(api.Cloud), "*"),
)
