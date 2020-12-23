package etcd

import (
	"coco-tool/config/provider"
	"context"
	"github.com/coreos/etcd/clientv3"
	"time"
)

var (
	DefaultEndpoint = []string{"127.0.0.1:2379"}
	DefaultTimeout  = time.Second * 5
)

type Etcd struct {
	Endpoint []string
	Timeout  time.Duration
	Username string
	Password string
	kv       clientv3.KV
}

func init() {
	defaultProvider, err := NewDefaultProvider()
	if err != nil {
		panic(err)
	}
	provider.Register("etcd", defaultProvider)
}

func NewDefaultProvider(endPoint ...string) (*Etcd, error) {
	addr := DefaultEndpoint
	if len(endPoint) > 0 {
		addr = endPoint
	}
	etcd := &Etcd{
		Endpoint: addr,
		Timeout:  DefaultTimeout,
	}
	err := etcd.init()
	return etcd, err
}

func NewProvider(username, password string, timeout time.Duration, endPoint ...string) (*Etcd, error) {

	etcd := &Etcd{
		Endpoint: endPoint,
		Timeout:  timeout,
		Username: username,
		Password: password,
	}
	err := etcd.init()
	return etcd, err
}

func (e *Etcd) init() error {
	cfg := clientv3.Config{
		Endpoints:   e.Endpoint,
		DialTimeout: e.Timeout,
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
