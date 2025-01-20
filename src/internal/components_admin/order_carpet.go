package components_admin

import (
	"service_api/internal/persist"
	"service_api/src/entitys"
	"service_api/src/util/consts"

	conf "bitbucket.org/klaraeng/package_config/conf"

	pb "bitbucket.org/klaraeng/package_proto/service_api/admin"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type AdminServiceOrderCarpet struct {
	pc    *conf.Params
	order persist.OrderToCarpetContext
	log   *log.Helper
}

func NewAdminOrderCarpetService(
	pc *conf.Params,
	logger log.Logger,
	order persist.OrderToCarpetContext) *AdminServiceOrderCarpet {
	return &AdminServiceOrderCarpet{
		pc:    pc,
		order: order,
		log:   log.NewHelper(logger),
	}
}

func (s *AdminServiceOrderCarpet) GetOrdersCarpet(req *pb.GetOrdersCarpetRequest) (*pb.GetOrdersCarpetReply, error) {
	var array []*pb.OrderCarpet
	var count int64
	count = s.order.FindListCount()

	if count < 1 {
		return nil, errors.New(500, "ORDERS", "Пусто")
	}

	offset := req.Point*req.Limit - req.Limit

	qb := s.order.FindList(req.Limit, offset)
	if qb == nil {
		return nil, errors.New(500, "ORDERS", "Ошибка в получении списка")
	}

	for _, el := range qb {
		if el == nil {
			continue
		}

		client := s.order.FindClientById(el.ClientId)
		if client == nil {
			client = &entitys.Client{
				Name:  "Не указано",
				Phone: "Не указано",
			}
		}

		article := s.order.FindProductById(int64(el.ProductId))
		array = append(array, &pb.OrderCarpet{
			Id:          int64(el.ID),
			Uuid:        el.Uuid,
			Name:        el.Name,
			Status:      consts.StatusOrder()[el.Status],
			Client:      client.Name,
			ClientPhone: client.Phone,
			Price:       el.Price,
			Height:      el.Height,
			Width:       el.Width,
			Count:       el.Count,
			Product:     article.Article,
			PriceUuid:   el.PriceUuid,
			PriceMid:    el.PriceMid,
			Date:        el.UpdateAt.String(),
		})
	}

	return &pb.GetOrdersCarpetReply{
		Success: true,
		Limit:   req.Limit,
		Point:   req.Point,
		Count:   count,
		Order:   array,
	}, nil
}

func (s *AdminServiceOrderCarpet) GetOrderCarpet(req *pb.GetOrderCarpetRequest) (*pb.GetOrderCarpetReply, error) {
	if len(req.Uuid) < 1 {
		return nil, errors.New(500, "ORDER_CARPET", "Не найдено")
	}

	el := s.order.FindByUUID(req.Uuid)
	if el == nil {
		return nil, errors.New(500, "ORDER_CARPET", "Ошибка в получении заказа")
	}

	client := s.order.FindClientById(el.ClientId)
	if client == nil {
		client = &entitys.Client{
			Name:  "Не указано",
			Phone: "Не указано",
		}
	}

	article := s.order.FindProductById(int64(el.ProductId))
	return &pb.GetOrderCarpetReply{
		Success: true,
		Order: &pb.OrderCarpet{
			Id:          int64(el.ID),
			Uuid:        el.Uuid,
			Name:        el.Name,
			Status:      consts.StatusOrder()[el.Status],
			Client:      client.Name,
			ClientPhone: client.Phone,
			Price:       el.Price,
			Height:      el.Height,
			Width:       el.Width,
			Count:       el.Count,
			Product:     article.Article,
			PriceUuid:   el.PriceUuid,
			PriceMid:    el.PriceMid,
			Date:        el.UpdateAt.String(),
		},
	}, nil
}

func (s *AdminServiceOrderCarpet) SetOrderCarpet(req *pb.SetOrderCarpetRequest) (*pb.SetOrderCarpetReply, error) {
	if len(req.Uuid) < 1 {
		return nil, errors.New(500, "ORDER_CARPET", "Не найдено")
	}

	el := s.order.FindByUUID(req.Uuid)
	if el == nil {
		return nil, errors.New(500, "ORDER_CARPET", "Ошибка в получении заказа")
	}

	for i, _ := range consts.StatusOrder() {
		if consts.StatusOrder()[i] == req.Order.Status {
			el.Status = i
		}
	}

	s.order.Update(el)

	return &pb.SetOrderCarpetReply{
		Success: true,
	}, nil
}
