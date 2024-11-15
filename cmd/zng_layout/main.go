package main

import (
	"github.com/zngue/zng_app/app"
	"github.com/zngue/zng_app/config"
	"github.com/zngue/zng_app/config/nacos"
	"github.com/zngue/zng_app/config/option"
	"github.com/zngue/zng_layout/internal/conf"
	"os"
	"strconv"
)

func main() {
	//获取环境变量
	var (
		cfg          *conf.Bootstrap
		err          error
		oriHost      = "nacos.zngue.com"
		oriNamespace = "develop"
		httpPort     = 16666
	)
	var host = os.Getenv("HOST")
	if host == "" {
		host = oriHost
	}
	var namespace = os.Getenv("NAMESPACE")
	if namespace == "" {
		namespace = oriNamespace
	}
	var port = os.Getenv("PORT")
	if port != "" {
		oriPort, _ := strconv.Atoi(port)
		if oriPort > 0 {
			httpPort = oriPort
		}
	}
	if len(host) == 0 {
		panic("配置中心请设置环境变量 HOST")
	}
	if len(namespace) == 0 {
		panic("配置中心请设置环境变量 NAMESPACE")
	}
	err = option.NewOption(&cfg, &option.Option{
		GroupName: "zng_layout",
		NaFns: []nacos.Fn{
			nacos.DataWithLogLevel(nacos.INFO),
			nacos.DataWithAppendToStdout(true),
			nacos.DataWithHost(host),
		},
		CFns: []config.Fn{
			config.WithDataId("config.yaml"),
		},
		RegisterNaFn: func(fn *nacos.CenterOptions) (fnErr error) {
			fnErr = fn.RegisterInstance(&nacos.RegisterInstanceParam{
				Port:        cfg.App.Port,
				ClusterName: "zng_layout",
				ServiceName: "zng_layout",
				GroupName:   "zng_layout",
			})
			return
		},
	})
	if err != nil {
		panic(err)
	}
	err = app.NewAppRunner(int32(httpPort), func() (*app.App, func(), error) {
		cfg.App.Port = int32(httpPort)
		return initApp(cfg)
	})
	if err != nil {
		panic(err)
	}
}
