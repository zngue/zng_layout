package http

import (
	"github.com/google/wire"
	"github.com/zngue/zng_app/app"
	v1 "github.com/zngue/zng_layout/internal/http/v1"
)

var ProviderSet = wire.NewSet(
	app.NewRouter,
	NewHttp,
	NewCron,
	v1.NewRouter,
	NewHttpGroup,
	NewService,
	NewApiService,
)
