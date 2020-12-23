package provider

import (
	"context"
	"sync"
)

type KVProvider interface {
	Set(context.Context, string, string) error
	Get(context.Context, string) (string, error)
	Del(context.Context, string) error
	Keys(context.Context) ([]string, error)
}

func init()  {
	NewProvider()
}

var ProviderSrv = &Provider{}

type Provider struct {
	mu sync.Mutex
	CurrentProvider string
	Providers       map[string]KVProvider
}

func (p *Provider) SetCurrent(name string)  {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.CurrentProvider = name
}

func NewProvider()  {
	ProviderSrv = &Provider{
		CurrentProvider: "etcd",
		Providers: make(map[string]KVProvider),
	}
}

func Register(name string, provider KVProvider) {
	ProviderSrv.Providers[name] = provider
}

func GetProvider(name string) KVProvider {
	return ProviderSrv.Providers[name]
}
