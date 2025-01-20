package auth

import (
	"service_api/internal/cache"
	"service_api/internal/persist"

	conf "bitbucket.org/klaraeng/package_config/conf"
)

type Service struct {
	admin persist.AdminContext
	//client persist.ClientContext
	cache *cache.Service
	conf  *conf.Params
}

func NewAuthService(
	admin persist.AdminContext,
	//client persist.ClientContext,
	cache *cache.Service,
	conf *conf.Params) *Service {
	return &Service{
		admin: admin,
		//client: client,
		cache: cache,
		conf:  conf,
	}
}
