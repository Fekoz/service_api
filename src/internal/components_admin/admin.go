package components_admin

import (
	"service_api/internal/persist"

	conf "bitbucket.org/klaraeng/package_config/conf"

	pb "bitbucket.org/klaraeng/package_proto/service_api/admin"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type AdminServiceAdmin struct {
	pc    *conf.Params
	admin persist.AdminContext
	log   *log.Helper
}

func NewAdminAdminService(
	pc *conf.Params,
	logger log.Logger,
	admin persist.AdminContext) *AdminServiceAdmin {
	return &AdminServiceAdmin{
		pc:    pc,
		admin: admin,
		log:   log.NewHelper(logger),
	}
}

func (s *AdminServiceAdmin) GetAdmins(req *pb.GetAdminsRequest) (*pb.GetAdminsReply, error) {
	var array []*pb.AdminProfile
	var count int64
	count = s.admin.FindListCount()

	if count < 1 {
		return nil, errors.New(500, "ADMIN", "Пусто")
	}

	offset := req.Point*req.Limit - req.Limit

	qb := s.admin.FindList(req.Limit, offset)
	if qb == nil {
		return nil, errors.New(500, "ADMIN", "Ошибка в получении списка")
	}

	for _, el := range qb {
		if el == nil {
			continue
		}

		array = append(array, &pb.AdminProfile{
			Id:         int64(el.ID),
			Login:      el.Login,
			Pass:       el.Pass,
			Email:      el.Email,
			IsActive:   el.IsActive,
			Name:       el.Name,
			Protection: el.Protection,
			Role:       el.Role,
			Date:       el.UpdateAt.String(),
		})
	}

	return &pb.GetAdminsReply{
		Success: true,
		Limit:   req.Limit,
		Point:   req.Point,
		Count:   count,
		Admin:   array,
	}, nil
}

func (s *AdminServiceAdmin) SetAdmin(req *pb.SetAdminRequest) (*pb.SetAdminReply, error) {
	isUpdate := s.admin.SetByID(req.Admin.Login, req.Admin.Name, req.Admin.Email, req.Admin.Pass, req.Admin.IsActive, req.Admin.Role, req.Admin.Protection)
	if true == isUpdate {
		return &pb.SetAdminReply{Success: true}, nil
	}

	return nil, errors.New(500, "ADMIN", "Ошибка при обновлении или записи")
}
