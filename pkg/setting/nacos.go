package setting

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/sirupsen/logrus"
)

type Nacos struct {
	ConfigClient config_client.IConfigClient
	NamingClient naming_client.INamingClient
	NamespaceId  string
	Group        string
	DataId       string
}

func NewConfigNacos(c chan interface{}, NamespaceId string, Group string, DataId string, IpAddr string, port uint64) *Nacos {
	sc := []constant.ServerConfig{
		{
			IpAddr: IpAddr,
			Port:   port,
		},
	}

	cc := newClientConfig(NamespaceId)

	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	nacos := &Nacos{
		ConfigClient: configClient,
		NamingClient: namingClient,
		NamespaceId:  NamespaceId,
		Group:        Group,
		DataId:       DataId,
	}
	go func() {
		nacos.Refresh(c)
	}()

	content := nacos.GetConfig()

	logrus.Info(content)

	return nacos
}

func newClientConfig(NamespaceId string) constant.ClientConfig {
	cc := constant.ClientConfig{
		NamespaceId:         NamespaceId, //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}
	return cc
}

func (n *Nacos) Refresh(c chan interface{}) {
	_ = n.ConfigClient.ListenConfig(vo.ConfigParam{
		DataId: n.DataId,
		Group:  n.Group,
		OnChange: func(namespace, group, dataId, data string) {
			logrus.Info("Config is refreshing from nacos")
			c <- data
		},
	})
}

func (n *Nacos) GetConfig() string {
	config, err := n.ConfigClient.GetConfig(vo.ConfigParam{
		DataId: n.DataId,
		Group:  n.Group,
	})
	if err != nil {
		panic(err)
	}
	return config
}

func (n *Nacos) RegisterServiceInstance(param vo.RegisterInstanceParam) (bool, error) {
	return n.NamingClient.RegisterInstance(param)
}
