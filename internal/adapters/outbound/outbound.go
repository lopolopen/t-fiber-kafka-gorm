package outbound

import (
	"github.com/google/wire"
	"github.com/lopolopen/t-fiber-kafka-gorm/internal/adapters/outbound/pubs"
)

var ProviderSet = wire.NewSet(pubs.NewPub)
