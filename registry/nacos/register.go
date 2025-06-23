package nacos

import (
	creg "github.com/adgers/common/registry"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	kreg "github.com/go-kratos/kratos/v2/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
)

type Nacos struct {
	reg creg.Registry
}

func NewNacos() *Nacos {
	return &Nacos{}
}

func (n Nacos) Register(c *creg.RegistryConfig) kreg.Registrar {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(c.Nacos.Registry.Address, c.Nacos.Registry.Port),
	}
	cc := constant.ClientConfig{
		NamespaceId:         c.Nacos.Registry.Namespace, // namespace id
		TimeoutMs:           5000,                       //配置项自己定义
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

func (n Nacos) Discovery(c *creg.RegistryConfig) kreg.Discovery {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(c.Nacos.Registry.Address, c.Nacos.Registry.Port),
	}
	cc := constant.ClientConfig{
		NamespaceId:         c.Nacos.Registry.Namespace, // namespace id
		TimeoutMs:           5000,                       //配置项自己定义
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
