package stop_list

import (
	"context"
	"service_api/internal/cache"
	"service_api/internal/persist"
	"service_api/src/util"
	"service_api/src/util/consts"

	conf "bitbucket.org/klaraeng/package_config/conf"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/metadata"
)

type Service struct {
	sl    persist.StopListContext
	cache *cache.Service
	conf  *conf.Params
}

func NewStopListService(sl persist.StopListContext, conf *conf.Params, cache *cache.Service) *Service {
	if conf.IsStopList == false {
		return nil
	}
	return &Service{
		sl:    sl,
		conf:  conf,
		cache: cache,
	}
}

func (s *Service) GetCreateOrderSelection(ctx context.Context) error {
	p, err := util.GetClientIp(ctx)
	if err == nil {
		result := s.sl.GetSlLimitToday(p, consts.CreateOrderSelection)
		if result == true {
			return errors.New(500, "CLIENT_LIMIT", "Лимит запросов. Клиент внесен в стоп-лист.")
		}
		return nil
	}
	return errors.New(500, "CLIENT_IP", "Нет доступа")
}

func (s *Service) AppendToCreateOrderSelection(ctx context.Context) {
	p, err := util.GetClientIp(ctx)
	if err == nil {
		s.sl.CreateOrUpdateSl(p, consts.CreateOrderSelection)
	}
}

func (s *Service) GetIp(ctx context.Context) (string, error) {
	return util.GetClientIp(ctx)
}

func (s *Service) GetSession(ctx context.Context) (string, error) {
	var session string
	if md, ok := metadata.FromServerContext(ctx); ok {
		session = md.Get(consts.MetadataAdminAuth)
	}

	if len(session) > 0 {
		return session, nil
	}

	return "", errors.New(500, "SESSION", "Не указана сессия. Не доступно для использования")
}

func (s *Service) SetSession(ctx context.Context, session string) {
	ctx = metadata.AppendToClientContext(ctx, consts.MetadataAdminAuth, session)
}

func (s *Service) IsLock(ctx context.Context, service string) error {
	if true == s.cache.IsUnLock(ctx, service) {
		return nil
	}
	return errors.New(500, "CLIENT_LIMIT", "Лимит запросов.")
}

func (s *Service) AppendLock(ctx context.Context, service string) {
	s.cache.AppendLock(ctx, service)
}

func (s *Service) DropLock(ctx context.Context, service string) {
	s.cache.DropLock(ctx, service)
}
