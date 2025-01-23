package main

import (
	"github.com/zngue/zng_app"
	"github.com/zngue/zng_app/app"
	"github.com/zngue/zng_app/config"
	"github.com/zngue/zng_app/config/nacos"
	"github.com/zngue/zng_app/config/option"
	"github.com/zngue/zng_app/log"
	"github.com/zngue/zng_app/pkg"
	"github.com/zngue/zng_layout/internal/conf"
	"github.com/zngue/zng_layout/pkg/util"
	"os"
)

func main() {
	//获取环境变量
	var (
		cfg           *conf.Bootstrap
		err           error
		serviceName   = "zng_layout"
		defaultConfig *pkg.DefaultConfig
	)
	zng_app.AppName = serviceName
	zng_app.SyncLogger = true
	defaultConfig, err = pkg.NewConfig()
	defaultConfig.Host = "39.98.204.118"
	if err != nil {
		log.Errorf("load config err NewConfig err %v", err)
		panic(err)
	}
	if defaultConfig.LogLevel != "" {
		err = os.Setenv("LOG_LEVEL", defaultConfig.LogLevel)
		if err != nil {
			log.Errorf("load config err Setenv err %v", err)
			panic(err)
		}
	}
	log.NewLog(util.LogConfig())
	//开启日志文件
	err = option.NewOption(&cfg, &option.Option{
		GroupName: defaultConfig.Group,
		NaFns: []nacos.Fn{
			nacos.DataWithLogLevel(nacos.INFO),
			nacos.DataWithAppendToStdout(false),
			nacos.DataWithHost(defaultConfig.Host),
		},
		CFns: []config.Fn{
			config.WithDataId("config.yaml"),
		},
		RegisterNaFn: RegisterFn(serviceName, defaultConfig),
	})
	if err != nil {
		log.Errorf("load config err NewOption err %v", err)
		panic(err)
	}
	//设置http 端口
	cfg.App.Port = int32(defaultConfig.HttpPort)
	err = app.NewAppRunner(int32(defaultConfig.HttpPort), NewHttpRun(cfg))
	if err != nil {
		log.Errorf("load config err NewAppRunner err %v", err)
		panic(err)
	}
}
func RegisterFn(serviceName string, defaultConfig *pkg.DefaultConfig) option.Fn {
	return func(fn *nacos.CenterOptions) (err error) {
		err = fn.RegisterInstance(&nacos.RegisterInstanceParam{
			Port:        int32(defaultConfig.HttpPort),
			ClusterName: serviceName,
			ServiceName: serviceName,
			GroupName:   serviceName,
		})
		return
	}
}

func NewHttpRun(config *conf.Bootstrap) app.Fn {
	return func() (*app.App, func(), error) {
		return initApp(config)
	}
}
