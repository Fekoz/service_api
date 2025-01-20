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

type ServiceOrderToCarpet struct {
	pc      *conf.Params
	order   persist.OrderToCarpetContext
	client  persist.ClientContext
	product persist.ProductContext
	price   persist.PriceContext
	queue   *queue.Queue
	log     *log.Helper
	tmp     persist.MailTemplateContext
	img     persist.ImagesContext
}

func NewOrderToCarpetService(
	pc *conf.Params,
	order persist.OrderToCarpetContext,
	client persist.ClientContext,
	product persist.ProductContext,
	price persist.PriceContext,
	queue *queue.Queue,
	tmp persist.MailTemplateContext,
	img persist.ImagesContext,
	logger log.Logger) *ServiceOrderToCarpet {
	return &ServiceOrderToCarpet{
		pc:      pc,
		order:   order,
		client:  client,
		product: product,
		price:   price,
		queue:   queue,
		tmp:     tmp,
		img:     img,
		log:     log.NewHelper(logger),
	}
}

func (s *ServiceOrderToCarpet) supplyCreate(
	uid string,
	client *entitys.Client,
	product *entitys.Product,
	price *entitys.Price,
	count int64,
) {

	orderType := consts.TypeOrder()[consts.OrderTypeCommission]
	orderNumber := consts.OrderTypeCommission
	if price.ID != 0 || len(price.Mid) > 1 || len(price.Uuid) > 1 {
		orderType = consts.TypeOrder()[consts.OrderTypeCommissionFull]
		orderNumber = consts.OrderTypeCommissionFull
	}

	s.order.CreateOrder(&entitys.OrderToCarpet{
		Status:    consts.OrderStatusStart,
		Name:      orderType,
		ClientId:  client.ID,
		ProductId: uint(product.ID),
		PriceMid:  price.Mid,
		Price:     int64(price.Price),
		Width:     price.Width,
		Height:    price.Height,
		Count:     count,
		Uuid:      uid,
	})

	params := map[string]string{
		"{Order}":     uid,
		"{Count}":     strconv.FormatInt(count, 10),
		"{TypeOrder}": orderType,
		"{Price}":     strconv.FormatInt(int64(price.Price), 10),
	}

	_ = s.queue.Send(
		consts.QueueList()[consts.QueueEmail],
		order_utils.SenderCarpetEmailCreate(
			params,
			s.tmp.FindByID(orderNumber),
			client,
			product,
			price,
			s.img.FindToImagesFirstEntity(product.ID),
		),
		"text/plain",
	)

	_ = s.queue.Send(
		consts.QueueList()[consts.QueueTelegram],
		order_utils.SenderCarpetTelegramCreate(
			fmt.Sprintf(
				"Новый заказ %s [%s] для клиента %s / +%s",
				orderType,
				uid,
				client.Name,
				client.Phone,
			),
		),
		"text/plain",
	)

}

func (s *ServiceOrderToCarpet) Create(req *pb.CreateOrderToCarpetRequest) (*pb.CreateOrderToCarpetReply, error) {
	client, err := s.client.CreateOrUpdateClientReturnClient(req.Client)
	if err != nil {
		return nil, errors.New(500, "CLIENT_NOT_CREATE", "Произошла ошибка в создании клиента")
	}

	product := s.product.FindByArticle(req.Article)
	if product == nil || len(product.Article) < 1 || product.ID == 0 {
		return nil, errors.New(500, "PRODUCT_NOT_CORRECT", "Переданный ковер некорректен")
	}

	uid := uuid.NewString()

	price := &entitys.Price{}

	go s.supplyCreate(uid, client, product, price, 0)

	return &pb.CreateOrderToCarpetReply{
		Success: true,
		Status:  consts.OrderStatusStart,
		Uid:     uid,
	}, nil
}

func (s *ServiceOrderToCarpet) CreatePrice(req *pb.CreateOrderToCarpetPriceRequest) (*pb.CreateOrderToCarpetPriceReply, error) {
	client, err := s.client.CreateOrUpdateClientReturnClient(req.Client)
	if err != nil {
		return nil, errors.New(500, "CLIENT_NOT_CREATE", "Произошла ошибка в создании клиента")
	}

	product := s.product.FindByArticle(req.Article)
	if product == nil || len(product.Article) < 1 || product.ID == 0 {
		return nil, errors.New(500, "PRODUCT_NOT_CORRECT", "Переданный ковер некорректен")
	}

	counter := int64(0)
	price := &entitys.Price{
		Mid:  req.MidPrice,
		ID:   uint(req.PriceId),
		Uuid: req.UuidPrice,
	}

	for _, field := range s.price.OneToProductIdPrice(product.ID) {
		if req.PriceId == int64(field.ID) || req.MidPrice == field.Mid || req.UuidPrice == field.Uuid {
			counter += field.Count
			price.Count = counter
			price.ProductId = field.ProductId
			price.Price = field.Price
			price.Height = field.Height
			price.Width = field.Width
		}
	}

	if price.ID == 0 && len(price.Mid) < 1 && len(price.Uuid) < 1 {
		return nil, errors.New(500, "PRODUCT_PRICE_NOT_CORRECT", "Переданный размер ковра не доступен")
	}

	if req.Count >= price.Count {
		req.Count = price.Count
	}

	uid := uuid.NewString()

	go s.supplyCreate(uid, client, product, price, req.Count)

	return &pb.CreateOrderToCarpetPriceReply{
		Success: true,
		Status:  consts.OrderStatusStart,
		Uid:     uid,
	}, nil
}
