package cache

import (
	"context"
	"fmt"
	"service_api/src/dtos"
	"service_api/src/util"
	"service_api/src/util/consts"
)

func (c *Service) IsUnLock(ctx context.Context, service string) bool {
	req, err := util.GetClientIp(ctx)
	if err != nil {
		return false
	}
	name := fmt.Sprintf(consts.MaskCacheKey, c.conf.Cache.Lock.Name, service)

	key, errGen := util.GenKey[string](name, &req)
	if errGen != nil {
		return false
	}

	if !c.is(key) {
		return true
	}

	lock, _ := util.DecodeValue[*dtos.LockTransfer](c.read(key))
	if lock.Try < consts.MaxTry {
		return true
	}

	return false
}

func (c *Service) AppendLock(ctx context.Context, service string) {
	req, err := util.GetClientIp(ctx)
	if err != nil {
		return
	}
	name := fmt.Sprintf(consts.MaskCacheKey, c.conf.Cache.Lock.Name, service)

	key, err := util.GenKey[string](name, &req)

	if err != nil {
		return
	}

	if !c.is(key) {
		var lock *dtos.LockTransfer
		lock = &dtos.LockTransfer{
			Ip:      service,
			Methode: name,
			Try:     1,
		}
		valueNew, errNew := util.EncodeValue[*dtos.LockTransfer](&lock)
		if errNew == nil {
			c.append(key, valueNew, c.conf.Cache.Lock.Ttl)
			return
		}

		return
	}

	lock, err := util.DecodeValue[*dtos.LockTransfer](c.read(key))
	if err != nil {
		return
	}

	lock.Try++
	result, err := util.EncodeValue[*dtos.LockTransfer](&lock)
	if err != nil {
		return
	}

	c.rewrite(key, result, c.conf.Cache.Lock.Ttl)
}

func (c *Service) DropLock(ctx context.Context, service string) {
	req, err := util.GetClientIp(ctx)
	if err != nil {
		return
	}
	name := fmt.Sprintf(consts.MaskCacheKey, c.conf.Cache.Lock.Name, service)

	key, _ := util.GenKey[string](name, &req)
	c.drop(key)
}
