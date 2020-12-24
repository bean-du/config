package etcd

import (
	"coco-tool/config/conf"
	"coco-tool/config/provider"
	"context"
	"github.com/coreos/etcd/clientv3"
	"time"
)

var (
	DefaultEtcd = &Etcd{Endpoint: []string{"127.0.0.1:2379"}, Timeout: time.Second * 5}
)

type Etcd struct {
	Endpoint []string
	Timeout  time.Duration
	Username string
	Password string
	kv       clientv3.KV
}

type Options func(e *Etcd)

func NewEtcd(options ...Options) *Etcd {
	etcd := DefaultEtcd
	for _, opt := range options {
		opt(etcd)
	}
	if err := etcd.init(); err != nil {
		panic(err)
	}
	return etcd
}

func WithEndpoint(endPoint ...string) Options {
	return func(e *Etcd) {
		e.Endpoint = endPoint
	}
}

func WithTimeout(timeout time.Duration) Options  {
	return func(e *Etcd) {
		e.Timeout = timeout
	}
}

func WithUsernameAndPassword(username, password string) Options  {
	return func(e *Etcd) {
		e.Username = username
		e.Password = password
	}
}

func Init() {
	c := conf.Conf.Etcd
	etcd := NewEtcd(WithEndpoint(c.Endpoint...), WithTimeout(time.Duration(c.Timeout)), WithUsernameAndPassword(c.Username, c.Password))
	provider.Register("etcd", etcd)
}

func (e *Etcd) init() error {
	cfg := clientv3.Config{
		Endpoints:   e.Endpoint,
		DialTimeout: e.Timeout * time.Second,
		Username:    e.Username,
		Password:    e.Password,
	}
	client, err := clientv3.New(cfg)
	if err != nil {
		return err
	}
	e.kv = clientv3.NewKV(client)
	return nil
}

func (e *Etcd) Set(ctx context.Context, key string, value string) error {
	_, err := e.kv.Put(ctx, key, value)
	if err != nil {
		return err
	}
	return nil
}

func (e *Etcd) Get(ctx context.Context, key string) (string, error) {
	get, err := e.kv.Get(ctx, key)
	if err != nil {
		return "", err
	}
	return string(get.Kvs[0].Value), nil
}

func (e *Etcd) Del(ctx context.Context, key string) error {
	_, err := e.kv.Delete(ctx, key)
	if err != nil {
		return err
	}
	return nil
}

func (e *Etcd) Keys(ctx context.Context) ([]string, error) {
	get, err := e.kv.Get(ctx, "", clientv3.WithPrefix(), clientv3.WithKeysOnly())
	if err != nil {
		return nil, err
	}
	keys := make([]string, 0, 0)
	for _, kvs := range get.Kvs {
		keys = append(keys, string(kvs.Key))
	}
	return keys, nil
}
