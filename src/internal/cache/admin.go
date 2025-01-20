package cache

import (
	"errors"
	"fmt"
	"service_api/src/dtos"
	"service_api/src/util"
	"service_api/src/util/consts"
)

func (c *Service) IsAdminCache(session string, ip string) bool {
	data, err := util.DecodeValue[*dtos.AdminTransfer](c.read(fmt.Sprintf(consts.MaskCacheKey, c.conf.Cache.Admin.Name, session)))

	if data != nil && data.Ip == ip && err == nil {
		return true
	}

	return false
}

func (c *Service) GetAdminCache(session string) (*dtos.AdminTransfer, error) {
	data, err := util.DecodeValue[*dtos.AdminTransfer](c.read(fmt.Sprintf(consts.MaskCacheKey, c.conf.Cache.Admin.Name, session)))
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.New("Admin session is not found")
	}

	return data, nil
}

func (c *Service) SetAdminCache(admin *dtos.AdminTransfer) string {
	key := util.GenSession(admin)
	if len(key) == 0 {
		return ""
	}

	data, err := util.EncodeValue[*dtos.AdminTransfer](&admin)
	if err != nil {
		return ""
	}

	if len(data) == 0 {
		return ""
	}

	c.append(fmt.Sprintf(consts.MaskCacheKey, c.conf.Cache.Admin.Name, key), data, c.conf.Cache.Admin.Ttl)

	return key
}
