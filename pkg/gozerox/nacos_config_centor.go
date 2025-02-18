package gozerox

import (

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/conf"
)

type Cfg struct {
	Nacos NacosConfig
}

type NacosConfig struct {
	NamespaceId         string
	TimeoutMs           uint64
	NotLoadCacheAtStart bool
	LogDir              string
	CacheDir            string
	LogLevel            string
	Group               string
	IpAddr              string
	Port                uint64
	ContextPath         string
	Scheme              string
	DataId              string
}

type NacosDistributedConfigCentor[T any] struct {
	cfg NacosConfig
	configClient config_client.IConfigClient
	localPath string
	data T
}

func MustInitNacosDistributedConfigCentor[T any](cfg NacosConfig) *NacosDistributedConfigCentor[T] {
	clientConfig := constant.ClientConfig{
		NamespaceId:         cfg.NamespaceId,
		TimeoutMs:           cfg.TimeoutMs,
		NotLoadCacheAtStart: cfg.NotLoadCacheAtStart,
		LogDir:              cfg.LogDir,
		CacheDir:            cfg.CacheDir,
		LogLevel:            cfg.LogLevel,
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      cfg.IpAddr,
			ContextPath: cfg.ContextPath,
			Port:        cfg.Port,
			Scheme:      cfg.Scheme,
		},
	}

	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		logx.Error(err)
		panic(err)
	}
	return &NacosDistributedConfigCentor[T]{
		cfg: cfg,
		configClient: configClient,
	}
}

func (d *NacosDistributedConfigCentor[T]) SetLocalPath(localPath string) *NacosDistributedConfigCentor [T]{
	d.localPath = localPath
	return d
}

func (d *NacosDistributedConfigCentor[T]) MustReadInConfig(dst T)  {
	content, err := d.configClient.GetConfig(vo.ConfigParam{
		DataId: d.cfg.DataId,
		Group:  d.cfg.Group,
	})
	if err != nil {
		logx.Error(err)
		panic(err)
	}
	if content != "" {
		err := conf.LoadConfigFromYamlBytes([]byte(content), dst)
		if err != nil {
			logx.Error(err)
			panic(err)
		}
		logx.Info("load config on nacos")
	} else {
		conf.MustLoad(d.localPath, &dst)
	}
	d.data = dst
}


func (d *NacosDistributedConfigCentor[T]) ListenOnConfig(callback func (data T)) {
	err := d.configClient.ListenConfig(vo.ConfigParam{
    DataId: d.cfg.DataId,
    Group:  d.cfg.Group,
    OnChange: func(namespace, group, dataId, data string) {
			d.MustReadInConfig(d.data)
			callback(d.data)
		},
	})
	if err != nil {
		logx.Error(err)
		panic(err)
	}
}