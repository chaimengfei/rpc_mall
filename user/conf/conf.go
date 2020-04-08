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
				User string
			}
		}
		Redis struct {
			DataSource string
			Auth       string
		}
		Endpoints []string
		RpcServerConfig *rpc_config.ServiceConf
		IntegralRpc *rpc_config.ClientConf
	}
)

var (
	UserConf   = &Conf{}
	configFile = flag.String("f", "./user/conf/conf.json", "user config")
)

func init() {
	flag.Parse()
	config.LoadFile(*configFile)
	config.Scan(UserConf)
}
