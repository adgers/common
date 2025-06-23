package nacos

import (
	"github.com/adgers/common/registry"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	kreg "github.com/go-kratos/kratos/v2/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
)

func Register(c *registry.RegistryConfig_Nacos_Registry) kreg.Registrar {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(c.Address, c.Port),
	}
	cc := constant.ClientConfig{
		NamespaceId:         c.Namespace, // namespace id
		TimeoutMs:           5000,        //配置项自己定义
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		log.Panic(err)
	}
	return nacos.New(client)
}
