package http

import (
	"github.com/google/wire"
	v1 "github.com/zngue/zng_layout/internal/http/v1"
)

var ProviderSet = wire.NewSet(
	NewRouter,
	NewHttp,
	NewCron,
	v1.NewRouter,
	NewHttpGroup,
	NewService,
)
