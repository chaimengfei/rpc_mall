package conf

import (
	"flag"
	"github.com/micro/go-micro/config"
	rpc_config "github.com/yakaa/grpcx/config"
)

type (
	Conf struct {
		Mode  string `json:"Mode"`
		Port  string `json:"Port"`
		Mysql struct {
			DataSource string
			Table      struct {
				Order string
			}
		}
		Redis struct {
			DataSource string
			Auth       string
		}
		IntegralRpc *rpc_config.ClientConf
		RpcServerConfig *rpc_config.ServiceConf
	}
)

var (
	OrderConf   = &Conf{}
	configFile = flag.String("f", "./order/conf/conf.json", "order config")
)

func init() {
	flag.Parse()
	config.LoadFile(*configFile)
	config.Scan(OrderConf)
}
