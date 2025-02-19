package server

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewCronService,
	NewHttpService,
	NewHttpEngine,
	NewHttpGroup,
	NewCombine,
	NewV1Router,
)
