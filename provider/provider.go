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

func init() {
	NewProvider()
}

var (
	ProviderSrv     = &Provider{}
	DefaultProvider = "etcd"
)

type Provider struct {
	mu              sync.Mutex
	CurrentProvider string
	Providers       map[string]KVProvider
}

func (p *Provider) SetCurrent(name string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.CurrentProvider = name
}

func NewProvider() {
	ProviderSrv = &Provider{
		CurrentProvider: DefaultProvider,
		Providers:       make(map[string]KVProvider),
	}
}

func Register(name string, provider KVProvider) {
	ProviderSrv.Providers[name] = provider
}

func GetProvider(name string) KVProvider {
	ProviderSrv.mu.Lock()
	defer ProviderSrv.mu.Unlock()
	if p, ok := ProviderSrv.Providers[name]; ok {
		return p
	} else {
		return ProviderSrv.Providers[DefaultProvider]
	}
}
