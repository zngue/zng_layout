package main

import (
	"github.com/zngue/zng_app/app"
	"github.com/zngue/zng_app/config"
	"github.com/zngue/zng_app/config/nacos"
	"github.com/zngue/zng_app/config/option"
	"github.com/zngue/zng_layout/internal/conf"
	"os"
)

func main() {
	//获取环境变量
	var (
		cfg *conf.Bootstrap
		err error
	)
	var host = os.Getenv("HOST")
	var namespace = os.Getenv("NAMESPACE")
	if len(host) == 0 {
		panic("配置中心请设置环境变量 HOST")
	}
	if len(namespace) == 0 {
		panic("配置中心请设置环境变量 NAMESPACE")
	}
	err = option.NewOption(&cfg, &option.Option{
		GroupName: "test",
		NaFns: []nacos.Fn{
			nacos.DataWithLogLevel(nacos.INFO),
			nacos.DataWithAppendToStdout(true),
			nacos.DataWithHost(host),
		},
		CFns: []config.Fn{
			config.WithDataConfig("nacos", "config.yaml"),
		},
		RegisterNaFn: func(fn *nacos.CenterOptions) (fnErr error) {
			fnErr = fn.RegisterInstance(&nacos.RegisterInstanceParam{
				Port:        cfg.App.Port,
				ClusterName: "zng_app",
				ServiceName: "zng_app",
				GroupName:   "test",
			})
			return
		},
	})
	if err != nil {
		panic(err)
	}
	err = app.NewAppRunner(cfg.App.Port, func() (*app.App, func(), error) {
		return initApp(cfg)
	})
	if err != nil {
		panic(err)
	}
}
