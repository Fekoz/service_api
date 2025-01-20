package components_admin

import (
	"context"
	"service_api/internal/auth"
	"service_api/internal/stop_list"
	"service_api/src/util/consts"

	conf "bitbucket.org/klaraeng/package_config/conf"

	pb "bitbucket.org/klaraeng/package_proto/service_api/admin"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type AdminServiceAuth struct {
	pc       *conf.Params
	log      *log.Helper
	auth     *auth.Service
	stopList *stop_list.Service
}

func NewAdminAuthService(
	pc *conf.Params,
	logger log.Logger,
	auth *auth.Service,
	stopList *stop_list.Service,
) *AdminServiceAuth {
	return &AdminServiceAuth{
		pc:       pc,
		log:      log.NewHelper(logger),
		auth:     auth,
		stopList: stopList,
	}
}

func (s *AdminServiceAuth) IsAuth(ctx context.Context) error {
	ip, err := s.stopList.GetIp(ctx)
	if err != nil {
		return errors.New(500, "AUTH", "Нет IP")
	}

	session, err := s.stopList.GetSession(ctx)
	if err != nil {
		return errors.New(500, "AUTH", "Нет Сессии в Metadata")
	}

	isSuccess := s.auth.Is(session, ip)
	if true == isSuccess {
		return nil
	}

	return errors.New(500, "AUTH", "Сессия не актуальна")
}

func (s *AdminServiceAuth) Auth(ctx context.Context, req *pb.GetAuthRequest) (*pb.GetAuthReply, error) {
	ip, err := s.stopList.GetIp(ctx)
	if err != nil {
		return nil, errors.New(500, "AUTH", "Нет IP")
	}

	if nil != s.stopList.IsLock(ctx, consts.CreatAdminAuth) {
		return nil, errors.New(500, "AUTH", "Заблокировано.")
	}

	session, err := s.auth.Set(req.Login, req.Password, req.Email, ip)
	if err != nil || len(session) == 0 {
		s.stopList.AppendLock(ctx, consts.CreatAdminAuth)
		return nil, errors.New(500, "AUTH", "Не доступно")
	}

	s.stopList.DropLock(ctx, consts.CreatAdminAuth)
	s.stopList.SetSession(ctx, session)

	return &pb.GetAuthReply{
		Success: true,
		Session: session,
	}, nil
}

func (s *AdminServiceAuth) GetAdmin(ctx context.Context) (*pb.GetAdminReply, error) {
	session, err := s.stopList.GetSession(ctx)
	if err != nil {
		return nil, errors.New(500, "AUTH", "Нет Сессии в Metadata")
	}

	profile, err := s.auth.Get(session)
	if err != nil {
		return nil, errors.New(500, "AUTH", "Нет данных в Сессии")
	}

	return &pb.GetAdminReply{
		Id:      int64(profile.Id),
		Success: true,
		Email:   profile.Email,
		Login:   profile.Login,
		Ip:      profile.Ip,
	}, nil
}

func (s *AdminServiceAuth) GetAdminInformation(ctx context.Context) (*pb.GetAdminInformationReply, error) {
	session, err := s.stopList.GetSession(ctx)
	if err != nil {
		return nil, errors.New(500, "AUTH", "Нет Сессии в Metadata")
	}

	profile, err := s.auth.GetFull(session)
	if err != nil {
		return nil, errors.New(500, "AUTH", "Нет данных в Сессии")
	}

	return &pb.GetAdminInformationReply{
		Success: true,
		Admin: &pb.AdminProfile{
			Id:         int64(profile.ID),
			Login:      profile.Login,
			Email:      profile.Email,
			Name:       profile.Name,
			Pass:       profile.Pass,
			Role:       profile.Role,
			Protection: profile.Protection,
			IsActive:   profile.IsActive,
			Date:       profile.UpdateAt.String(),
		},
	}, nil
}

func (s *AdminServiceAuth) CacheClear() {
	s.auth.FlushAll()
}
