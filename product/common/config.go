package common

import (
	"github.com/asim/go-micro/plugins/config/source/consul/v4"
	"go-micro.dev/v4/config"
	"strconv"
)

func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	consulSrc := consul.NewSource(
		consul.WithAddress(host+":"+strconv.Itoa(int(port))),
		consul.WithPrefix(prefix),
		//表示可以不带前缀直接获取对应配置
		consul.StripPrefix(true),
	)
	conf, err := config.NewConfig()
	if err != nil {
		return conf, err
	}
	conf.Load(consulSrc)
	return conf, nil
}
