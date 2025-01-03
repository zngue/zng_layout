package server

import (
	"github.com/google/wire"
	"github.com/zngue/zng_layout/internal/server/http"
)

var ProviderSet = wire.NewSet(
	NewCron,
	http.ProviderSet,
	ProviderSetRouter,
)
