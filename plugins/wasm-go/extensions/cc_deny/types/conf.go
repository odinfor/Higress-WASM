package types

import (
	"gopkg.in/yaml.v3"
	"os"
)

type (
	conf struct {
		Redis      redis `yaml:"redis"`
		HeaderRule rule  `yaml:"HeaderRule"`
		CookieRule rule  `yaml:"CookieRule"`
	}

	rule struct {
		// 规则鉴定标签,支持两种调用方识别方式：
		// header: 从特定 header 识别
		// cookie: 从特定 cookie 识别
		Tag string `yaml:"tag"`

		// 每个调用方每秒最多调用次数
		QPS int `yaml:"qps"`

		// 每个调用方每分钟最多调用次数
		QPM int `yaml:"qpm"`

		// 每个调用方每天最多调用次数
		QPD int `yaml:"qpd"`

		// 超过限制后将该调用方屏蔽，不可访问。(单位秒)
		BlockSec int `yaml:"block_seconds"`
	}

	redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	}
)

func NewConfDo() ConfDo {
	return &conf{}
}

type ConfDo interface {
	HeaderConf() *rule
	CookieConf() *rule
	RedisConf() redis
}

func InitConf(filePath string) {
	if file, err := os.ReadFile(filePath); err != nil {
		panic("Open configuration file found error. err: " + err.Error())
	} else {
		if err = yaml.Unmarshal(file, NewConfDo()); err != nil {
			panic("Parsing configuration file found error. err: " + err.Error())
		}
	}
}

func (c *conf) HeaderConf() *rule {
	return &c.HeaderRule
}

func (c *conf) CookieConf() *rule {
	return &c.CookieRule
}

func (c *conf) RedisConf() redis {
	return c.Redis
}
