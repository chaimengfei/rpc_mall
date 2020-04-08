package conf

import (
	"flag"
	"github.com/micro/go-micro/config"
	rpc_config "github.com/yakaa/grpcx/config"
)

type Conf struct {
	RpcServerConfig *rpc_config.ServiceConf
	Mode            string `json:"mode"`
	Mysql           struct {
		DataSource string
		Table      struct {
			Integral string
		}
	}
	Redis struct {
		DataSource string
		Auth       string
	}
	RabbitMq struct {
		DataSource  string
		VirtualHost string
		QueueName   string
	}
}

var (
	IntegralConf   = &Conf{}
	configFile = flag.String("f", "./integral/conf/conf.json", "integral config")
)

func init() {
	flag.Parse()
	config.LoadFile(*configFile)
	config.Scan(IntegralConf)
}
