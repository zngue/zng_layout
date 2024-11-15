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
		configGroup  = "common"
		serviceName  = "zng_layout"
	)
	var host = os.Getenv("HOST")
	if host == "" {
		host = oriHost
	}
	//设置配置文件默认值
	var dbGroupName = os.Getenv("DB_GROUP")
	if dbGroupName != "" {
		configGroup = dbGroupName
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
		GroupName: configGroup,
		NaFns: []nacos.Fn{
			nacos.DataWithLogLevel(nacos.INFO),
			nacos.DataWithAppendToStdout(false),
			nacos.DataWithHost(host),
		},
		CFns: []config.Fn{
			config.WithDataId("config.yaml"),
		},
		RegisterNaFn: func(fn *nacos.CenterOptions) (fnErr error) {
			cfg.App.Port = int32(httpPort)
			fnErr = fn.RegisterInstance(&nacos.RegisterInstanceParam{
				Port:        cfg.App.Port,
				ClusterName: serviceName,
				ServiceName: serviceName,
				GroupName:   serviceName,
			})
			return
		},
	})
	if err != nil {
		panic(err)
	}
	err = app.NewAppRunner(int32(httpPort), func() (*app.App, func(), error) {
		return initApp(cfg)
	})
	if err != nil {
		panic(err)
	}
}
