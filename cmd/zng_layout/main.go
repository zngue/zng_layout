package main

import (
	"fmt"
	"github.com/zngue/zng_app"
	"github.com/zngue/zng_app/app"
	"github.com/zngue/zng_app/config"
	"github.com/zngue/zng_app/config/nacos"
	"github.com/zngue/zng_app/config/option"
	"github.com/zngue/zng_app/log"
	"github.com/zngue/zng_layout/internal/conf"
	"github.com/zngue/zng_layout/pkg/util"
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
		httpPort     = 16667
		configGroup  = "common"
		serviceName  = "zng_layout"
	)
	zng_app.AppName = serviceName
	zng_app.SyncLogger = true
	log.NewLog(util.LogConfig())
	var host = os.Getenv("NACOS_HOST")
	if host == "" {
		host = oriHost
	}
	log.Info(fmt.Sprintf("nacos host %s", host))
	//设置配置文件默认值
	var dbGroupName = os.Getenv("DB_GROUP")
	if dbGroupName != "" {
		configGroup = dbGroupName
	}
	log.Info(fmt.Sprintf("db group name %s", dbGroupName))
	var namespace = os.Getenv("NACOS_NAMESPACE")
	if namespace == "" {
		namespace = oriNamespace
	}
	log.Info(fmt.Sprintf("nacos namespace %s", namespace))
	var port = os.Getenv("HTTP_PORT")
	if port != "" {
		oriPort, _ := strconv.Atoi(port)
		if oriPort > 0 {
			httpPort = oriPort
		}
	}
	log.Info(fmt.Sprintf("http port %d", httpPort))
	if len(host) == 0 {
		panic("配置中心请设置环境变量 HOST")
	}
	if len(namespace) == 0 {
		panic("配置中心请设置环境变量 NAMESPACE")
	}
	//开启日志文件
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
		log.Errorf("load config err NewOption err %v", err)
		panic(err)
	}

	err = app.NewAppRunner(int32(httpPort), func() (*app.App, func(), error) {
		return initApp(cfg)
	})
	if err != nil {
		log.Errorf("load config err NewAppRunner err %v", err)
		panic(err)
	}
}
