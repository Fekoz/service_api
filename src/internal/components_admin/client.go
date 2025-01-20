package components_admin

import (
	"service_api/internal/persist"

	conf "bitbucket.org/klaraeng/package_config/conf"

	pb "bitbucket.org/klaraeng/package_proto/service_api/admin"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type AdminServiceClient struct {
	pc     *conf.Params
	client persist.ClientContext
	log    *log.Helper
}

func NewAdminClientService(
	pc *conf.Params,
	logger log.Logger,
	client persist.ClientContext) *AdminServiceClient {
	return &AdminServiceClient{
		pc:     pc,
		client: client,
		log:    log.NewHelper(logger),
	}
}

func (s *AdminServiceClient) GetClients(req *pb.GetClientsRequest) (*pb.GetClientsReply, error) {
	var array []*pb.Personal
	var count int64
	count = s.client.FindListCount()

	if count < 1 {
		return nil, errors.New(500, "CLIENT", "Пусто")
	}

	offset := req.Point*req.Limit - req.Limit

	qb := s.client.FindList(req.Limit, offset)
	if qb == nil {
		return nil, errors.New(500, "CLIENT", "Ошибка в получении списка")
	}

	for _, el := range qb {
		if el == nil {
			continue
		}
		array = append(array, &pb.Personal{
			Id:    int64(el.ID),
			Name:  el.Name,
			Email: el.Email,
			Phone: el.Phone,
		})
	}

	return &pb.GetClientsReply{
		Success: true,
		Limit:   req.Limit,
		Point:   req.Point,
		Count:   count,
		Client:  array,
	}, nil
}
