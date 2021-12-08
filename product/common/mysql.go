package common

import "go-micro.dev/v4/config"

type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

func GetMysqlConfigFromConsul(config config.Config, path ...string) *MysqlConfig {
	mysqlConf := &MysqlConfig{}
	config.Get(path...).Scan(mysqlConf)
	return mysqlConf
}
