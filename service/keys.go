package service

import (
	"coco-tool/config/model/entity"
	"coco-tool/config/provider"
	"coco-tool/config/repostory"
	"context"
	"fmt"
	"time"
)

var KeysSrv = &keysSrv{}

type keysSrv struct{}

func (k *keysSrv) Keys(ctx context.Context) ([]string, error) {
	keys := make([]string, 0, 0)
	etcd := provider.GetProvider(provider.ProviderSrv.CurrentProvider)
	etcdKeys, err := etcd.Keys(ctx)
	if err != nil {
		return nil, err
	}
	keyMap := make(map[string]struct{})
	for _, key := range etcdKeys {
		keyMap[key] = struct{}{}
	}
	// 从数据库取keys
	dbKeys, err := repostory.RecordRepo.Keys()
	for _, key := range dbKeys {
		keyMap[key] = struct{}{}
	}

	for k := range keyMap {
		keys = append(keys, k)
	}
	return keys, nil
}

func (k *keysSrv) KeyDetails(_ context.Context, key string) ([]*entity.Record, error) {
	return repostory.RecordRepo.Details(key)
}

func (k *keysSrv) Set(ctx context.Context, key, value string) error {
	record := &entity.Record{
		Key:       key,
		Value:     value,
		Version:   versionGenerate(),
		Pointer:   "yes",
		CreatedAt: time.Now().String(),
	}

	etcd := provider.GetProvider(provider.ProviderSrv.CurrentProvider)
	if err := etcd.Set(ctx, key, value); err != nil {
		return err
	}
	return repostory.RecordRepo.Set(record)
}

func (k *keysSrv) Get(ctx context.Context, key, version string) (*entity.Record, error) {
	if version == "" {
		get, err := provider.GetProvider(provider.ProviderSrv.CurrentProvider).Get(ctx, key)
		if err != nil {
			return nil, err
		}
		record := &entity.Record{
			Key:   key,
			Value: get,
		}
		return record, nil
	}
	return repostory.RecordRepo.Get(key, version)
}

func (k *keysSrv) Apply(ctx context.Context, key, version, value string) error {
	if err := provider.GetProvider(provider.ProviderSrv.CurrentProvider).Set(ctx, key, value); err != nil {
		return err
	}
	return repostory.RecordRepo.Apply(key, version)
}

func (k *keysSrv) Del(ctx context.Context, key, version string) error {
	if version == "" {
		if err := provider.GetProvider(provider.ProviderSrv.CurrentProvider).Del(ctx, key); err != nil {
			return err
		}
	}
	return repostory.RecordRepo.Del(key, version)
}

func versionGenerate() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}
