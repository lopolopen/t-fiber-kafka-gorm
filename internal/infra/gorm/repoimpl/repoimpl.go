package repoimpl

import (
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/domain/repo"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewUserRepo,
	wire.Bind(new(repo.UserRepo), new(*UserRepo)),
)
