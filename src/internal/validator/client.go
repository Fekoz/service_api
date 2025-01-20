package validator

import (
	"service_api/src/util"
	"service_api/src/util/consts"

	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
	"github.com/go-kratos/kratos/v2/errors"
)

type ServiceClient struct {
}

func NewValidatorClientService() *ServiceClient {
	return &ServiceClient{}
}

func (s *ServiceClient) GetViewcase(req *pb.GetViewcaseRequest) (bool, error) {
	if len(req.Uuid) < 1 {
		return false, errors.New(500, "UUID_IS_NULL", "Не указано значение UUID")
	}
	return true, nil
}

func (s *ServiceClient) GetShowcase(req *pb.GetShowcaseRequest) (bool, error) {
	if len(req.Uuid) < 1 {
		return false, errors.New(500, "UUID_IS_NULL", "Не указано значение UUID")
	}
	return true, nil
}

func (s *ServiceClient) GetCarpet(req *pb.GetCarpetRequest) (bool, error) {
	if len(req.Article) < 1 && len(req.Uuid) < 1 && req.Id == 0 {
		return false, errors.New(500, "REQUEST_IS_NULL", "Не указаны значения ARTICLE, UUID, ID")
	}
	return true, nil
}

func (s *ServiceClient) GetCarpets(req *pb.GetCarpetsRequest) *pb.GetCarpetsRequest {
	if req.Point < consts.DefaultPoint {
		req.Point = consts.DefaultPoint
	}

	if req.Limit < consts.MinLimit {
		req.Limit = consts.MinLimit
	}

	if req.Limit > consts.MaxLimit {
		req.Limit = consts.MaxLimit
	}

	return req
}

func (s *ServiceClient) CreateOrderLanding(req *pb.CreateOrderLandingRequest) (bool, error) {
	isClient := util.ValidateClient(req.Client)
	if isClient != true {
		return false, errors.New(500, "REQUEST_CLIENT_EMPTY", "Блок клиента некорректно заполнен")
	}

	if req.Size == nil {
		return false, errors.New(500, "REQUEST_SIZE_EMPTY", "Не указан блок SIZE")
	}
	if req.Size.Width <= 1 || req.Size.Height <= 1 {
		return false, errors.New(500, "REQUEST_SIZE_EMPTY", "Не указано значение WIDTH, HEIGHT")
	}

	return true, nil
}

func (s *ServiceClient) GetCatalog(req *pb.GetCatalogRequest) *pb.GetCatalogRequest {
	if req.Point < consts.DefaultPoint {
		req.Point = consts.DefaultPoint
	}

	if req.Limit < consts.MinLimit {
		req.Limit = consts.MinLimit
	}

	if req.Limit > consts.MaxLimit {
		req.Limit = consts.MaxLimit
	}

	if req.Price == nil {
		req.Price = &pb.GetCatalogRequest_PriceFind{}
	}

	if req.Product == nil {
		req.Product = &pb.GetCatalogRequest_ProductFind{}
	}

	return req
}

func (s *ServiceClient) CreateOrderToCarpet(req *pb.CreateOrderToCarpetRequest) (bool, error) {
	isClient := util.ValidateClient(req.Client)
	if isClient != true {
		return false, errors.New(500, "REQUEST_CLIENT_EMPTY", "Блок клиента некорректно заполнен")
	}

	if len(req.Article) < 1 {
		return false, errors.New(500, "REQUEST_ARTICLE_EMPTY", "Не указан необходимый артикул товара")
	}

	return true, nil
}

func (s *ServiceClient) CreateOrderToCarpetPrice(req *pb.CreateOrderToCarpetPriceRequest) (bool, error) {
	isClient := util.ValidateClient(req.Client)
	if isClient != true {
		return false, errors.New(500, "REQUEST_CLIENT_EMPTY", "Блок клиента некорректно заполнен")
	}

	if len(req.Article) < 1 {
		return false, errors.New(500, "REQUEST_ARTICLE_EMPTY", "Не указан необходимый артикул товара")
	}

	if req.Count < 1 {
		return false, errors.New(500, "REQUEST_COUNT_EMPTY", "Не указано количество товара")
	}

	return true, nil
}
