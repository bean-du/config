package consul

import (
	"coco-tool/config/conf"
	"coco-tool/config/provider"
	"context"
	"github.com/hashicorp/consul/api"
)

type Consul struct {
	Conf *api.Config
	User string
	Pwd  string
	Kv   *api.KV
}

func Init() {
	c := conf.Conf.Consul
	consul := NewConsul(WithConfig(newConsulConfig(c)), WithUserAndPwd("", ""))
	provider.Register("consul", consul)
}

func WithConfig(conf *api.Config) Option {
	return func(c *Consul) {
		c.Conf = conf
	}
}

func WithUserAndPwd(user, pwd string) Option {
	return func(c *Consul) {
		c.User = user
		c.Pwd = pwd
	}
}

func newConsulConfig(c struct{}) *api.Config {
	return &api.Config{}
}

func (c *Consul) Set(ctx context.Context, key string, val string) error {
	panic("implement me")
}

func (c *Consul) Get(ctx context.Context, key string) (string, error) {
	panic("implement me")
}

func (c *Consul) Del(ctx context.Context, key string) error {
	panic("implement me")
}

func (c *Consul) Keys(ctx context.Context) ([]string, error) {
	panic("implement me")
}

type Option func(c *Consul)

var DefaultConsul = &Consul{Conf: api.DefaultConfig()}

func NewConsul(option ...Option) *Consul {
	c := DefaultConsul
	for _, opt := range option {
		opt(c)
	}
	c.init()
	return c
}

func (c *Consul) init() {
	client, err := api.NewClient(c.Conf)
	if err != nil {
		panic(err)
	}
	c.Kv = client.KV()
}
