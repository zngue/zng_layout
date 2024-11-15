package model

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewDB,
	NewRedis,
	NewTest,
	NewUser,
	NewMember,
)
