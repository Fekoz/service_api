package components

import (
	"fmt"
	"service_api/internal/persist"
	"service_api/src/entitys"
	"service_api/src/util/consts"
	"service_api/src/util/order_utils"
	"strconv"

	queue "bitbucket.org/klaraeng/package_queue/ampq"

	conf "bitbucket.org/klaraeng/package_config/conf"

	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type ServiceOrderSelection struct {
	pc     *conf.Params
	order  persist.OrderContext
	client persist.ClientContext
	log    *log.Helper
	queue  *queue.Queue
	tmp    persist.MailTemplateContext
}

func NewOrderSelectionService(
	pc *conf.Params,
	order persist.OrderContext,
	client persist.ClientContext,
	logger log.Logger,
	queue *queue.Queue,
	tmp persist.MailTemplateContext,
) *ServiceOrderSelection {
	return &ServiceOrderSelection{
		pc:     pc,
		order:  order,
		client: client,
		log:    log.NewHelper(logger),
		queue:  queue,
		tmp:    tmp,
	}
}

func (s *ServiceOrderSelection) supplyCreate(client *entitys.Client, orderId uint, req *pb.CreateOrderLandingRequest, uid string) {
	s.order.CreateOrderFields(order_utils.OrderFieldCreate(orderId, req))

	params := map[string]string{
		"{OrderId}": strconv.Itoa(int(orderId)),
		"{Form}":    pb.CARPET_FORM_name[int32(req.Form)],
		"{Design}":  pb.CARPET_DESIGN_name[int32(req.Design)],
		"{Color}":   pb.CARPET_COLOR_name[int32(req.Color)],
		"{Width}":   strconv.FormatInt(req.Size.Width, 10),
		"{Height}":  strconv.FormatInt(req.Size.Height, 10),
	}

	_ = s.queue.Send(
		consts.QueueList()[consts.QueueEmail],
		order_utils.SenderCarpetEmailCreate(
			params,
			s.tmp.FindByID(1),
			client,
			&entitys.Product{},
			&entitys.Price{},
			&entitys.Images{},
		),
		"text/plain",
	)

	_ = s.queue.Send(
		consts.QueueList()[consts.QueueTelegram],
		order_utils.SenderCarpetTelegramCreate(
			fmt.Sprintf(
				"Новый заказ [(%d) / %s] на подбор ковра для клиента %s / +%s",
				orderId,
				uid,
				client.Name,
				client.Phone,
			),
		),
		"text/plain",
	)
}

func (s *ServiceOrderSelection) Create(req *pb.CreateOrderLandingRequest) (*pb.CreateOrderLandingReply, error) {
	client, err := s.client.CreateOrUpdateClientReturnClient(req.Client)
	if err != nil {
		return nil, errors.New(500, "CLIENT_NOT_CREATE", "Произошла ошибка в создании клиента")
	}

	newUuid := uuid.NewString()

	order := &entitys.Order{
		Status:   consts.OrderStatusStart,
		Type:     consts.OrderTypeSelection,
		Name:     consts.TypeOrder()[consts.OrderTypeSelection],
		ClientId: client.ID,
		Uuid:     newUuid,
	}

	codeId := consts.CARPET_PRESENT_ORDER_SELECTION
	id, err := s.order.GetPresentWithCode(codeId)
	if err == nil && id != 0 {
		order.PresentsId = id
	}

	orderId, err := s.order.CreateOrder(order)
	if err != nil {
		return nil, errors.New(500, "ORDER_NOT_CREATE", "Произошла ошибка в создании заказа")
	}

	go s.supplyCreate(client, orderId, req, newUuid)

	return &pb.CreateOrderLandingReply{
		Success: true,
		Uid:     newUuid,
		Status:  consts.OrderStatusStart,
	}, nil
}
