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

type AdminServiceOrder struct {
	pc    *conf.Params
	order persist.OrderContext
	log   *log.Helper
}

func NewAdminOrderService(
	pc *conf.Params,
	logger log.Logger,
	order persist.OrderContext) *AdminServiceOrder {
	return &AdminServiceOrder{
		pc:    pc,
		order: order,
		log:   log.NewHelper(logger),
	}
}

func (s *AdminServiceOrder) GetOrders(req *pb.GetOrdersRequest) (*pb.GetOrdersReply, error) {
	var array []*pb.Order
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

		var arrayField []*pb.Order_Field
		qbFd := s.order.FindOrderFieldById(el.ID)

		if qbFd != nil {
			for _, fd := range qbFd {
				arrayField = append(arrayField, &pb.Order_Field{
					Key:   fd.Key,
					Value: fd.Value,
				})
			}
		}

		client := s.order.FindClientById(el.ClientId)
		if client == nil {
			client = &entitys.Client{
				Name:  "Не указано",
				Phone: "Не указано",
			}
		}

		present := s.order.FindPresentById(el.PresentsId)
		if present == nil {
			present = &entitys.Presents{
				Name:  "Не указано",
				Value: "Не указано",
			}
		}

		array = append(array, &pb.Order{
			Id:          int64(el.ID),
			Uuid:        el.Uuid,
			Name:        el.Name,
			Status:      consts.StatusOrder()[el.Status],
			Type:        consts.TypeOrder()[el.Type],
			Field:       arrayField,
			Client:      client.Name,
			ClientPhone: client.Phone,
			Present:     present.Value,
			PresentName: present.Value,
			Date:        el.UpdateAt.String(),
		})
	}

	return &pb.GetOrdersReply{
		Success: true,
		Limit:   req.Limit,
		Point:   req.Point,
		Count:   count,
		Order:   array,
	}, nil
}

func (s *AdminServiceOrder) GetOrder(req *pb.GetOrderRequest) (*pb.GetOrderReply, error) {
	if len(req.Uuid) < 1 {
		return nil, errors.New(500, "ORDER", "Не найдено")
	}

	el := s.order.FindByUUID(req.Uuid)
	if el == nil {
		return nil, errors.New(500, "ORDER", "Ошибка в получении заказа")
	}

	var arrayField []*pb.Order_Field
	qbFd := s.order.FindOrderFieldById(el.ID)

	if qbFd != nil {
		for _, fd := range qbFd {
			arrayField = append(arrayField, &pb.Order_Field{
				Key:   fd.Key,
				Value: fd.Value,
			})
		}
	}

	client := s.order.FindClientById(el.ClientId)
	if client == nil {
		client = &entitys.Client{
			Name:  "Не указано",
			Phone: "Не указано",
		}
	}

	present := s.order.FindPresentById(el.PresentsId)
	if present == nil {
		present = &entitys.Presents{
			Name:  "Не указано",
			Value: "Не указано",
		}
	}

	return &pb.GetOrderReply{
		Success: true,
		Order: &pb.Order{
			Id:          int64(el.ID),
			Uuid:        el.Uuid,
			Name:        el.Name,
			Status:      consts.StatusOrder()[el.Status],
			Type:        consts.TypeOrder()[el.Type],
			Field:       arrayField,
			Client:      client.Name,
			ClientPhone: client.Phone,
			Present:     present.Value,
			PresentName: present.Value,
			Date:        el.UpdateAt.String(),
		},
	}, nil
}

func (s *AdminServiceOrder) SetOrder(req *pb.SetOrderRequest) (*pb.SetOrderReply, error) {
	if len(req.Uuid) < 1 {
		return nil, errors.New(500, "ORDER", "Не найдено")
	}

	el := s.order.FindByUUID(req.Uuid)
	if el == nil {
		return nil, errors.New(500, "ORDER", "Ошибка в получении заказа")
	}

	for i, _ := range consts.StatusOrder() {
		if consts.StatusOrder()[i] == req.Order.Status {
			el.Status = i
		}
	}

	for i, _ := range consts.TypeOrder() {
		if consts.TypeOrder()[i] == req.Order.Type {
			el.Type = i
		}
	}

	s.order.Update(el)

	return &pb.SetOrderReply{
		Success: true,
	}, nil
}
