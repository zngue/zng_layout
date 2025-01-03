package http

import (
	"github.com/google/wire"
	v1 "github.com/zngue/zng_layout/internal/server/http/v1"
)

var ProviderSet = wire.NewSet(
	NewHttp,
	v1.NewV1Router,
	v1.NewLoginV1Router,
	NewHttpGroup,
	NewService,
)
