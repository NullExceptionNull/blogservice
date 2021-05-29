package setting

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/sirupsen/logrus"
)

type Nacos struct {
	Client      config_client.IConfigClient
	NamespaceId string
	Group       string
	DataId      string
}

func NewNacos(c chan interface{}, NamespaceId string, Group string, DataId string, IpAddr string, port uint64) *Nacos {
	sc := []constant.ServerConfig{
		{
			IpAddr: IpAddr,
			Port:   port,
		},
	}

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
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}

	nacos := &Nacos{
		Client:      client,
		NamespaceId: NamespaceId,
		Group:       Group,
		DataId:      DataId,
	}
	go func() {
		nacos.Refresh(c)
	}()

	content := nacos.GetConfig()

	logrus.Info(content)

	return nacos
}

func (n *Nacos) Refresh(c chan interface{}) {
	_ = n.Client.ListenConfig(vo.ConfigParam{
		DataId: n.DataId,
		Group:  n.Group,
		OnChange: func(namespace, group, dataId, data string) {
			logrus.Info("Config is refreshing from nacos")
			c <- data
		},
	})
}

func (n *Nacos) GetConfig() string {
	config, err := n.Client.GetConfig(vo.ConfigParam{
		DataId: n.DataId,
		Group:  n.Group,
	})
	if err != nil {
		panic(err)
	}
	return config
}
