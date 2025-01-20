package validator

import (
	"service_api/src/util/consts"

	pb "bitbucket.org/klaraeng/package_proto/service_api/admin"
	"github.com/go-kratos/kratos/v2/errors"
)

type ServiceAdmin struct {
}

func NewValidatorAdminService() *ServiceAdmin {
	return &ServiceAdmin{}
}

func (s *ServiceAdmin) GetShowcases(req *pb.GetShowcasesRequest) *pb.GetShowcasesRequest {
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

func (s *ServiceAdmin) GetCatalog(req *pb.GetCatalogRequest) *pb.GetCatalogRequest {
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

func (s *ServiceAdmin) GetMailTemplates(req *pb.GetMailTemplatesRequest) *pb.GetMailTemplatesRequest {
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

func (s *ServiceAdmin) SetMailTemplate(req *pb.SetMailTemplateRequest) error {
	if req.Id == 0 || len(req.Title) < 1 || len(req.Message) < 1 {
		return errors.New(500, "REQUEST_IS_NULL", "Не указаны значения ID, TITLE, MESSAGE")
	}

	return nil
}

func (s *ServiceAdmin) GetMarketSequences(req *pb.GetMarketSequencesRequest) *pb.GetMarketSequencesRequest {
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

func (s *ServiceAdmin) SetMarketSequence(req *pb.SetMarketSequenceRequest) error {
	if len(req.Mid) < 1 {
		return errors.New(500, "REQUEST_IS_NULL", "Не указаны значения MID")
	}

	return nil
}

func (s *ServiceAdmin) GetClients(req *pb.GetClientsRequest) *pb.GetClientsRequest {
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

func (s *ServiceAdmin) GetAdmins(req *pb.GetAdminsRequest) *pb.GetAdminsRequest {
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

func (s *ServiceAdmin) SetAdmin(req *pb.SetAdminRequest) error {
	if len(req.Admin.Login) < 1 || len(req.Admin.Name) < 1 || len(req.Admin.Email) < 1 || req.Admin.Role == 0 {
		return errors.New(500, "REQUEST_IS_NULL", "Не указаны значения для AdminProfile")
	}

	return nil
}

func (s *ServiceAdmin) GetStopLists(req *pb.GetStopListsRequest) *pb.GetStopListsRequest {
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

func (s *ServiceAdmin) GetOrders(req *pb.GetOrdersRequest) *pb.GetOrdersRequest {
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

func (s *ServiceAdmin) GetOrdersCarpet(req *pb.GetOrdersCarpetRequest) *pb.GetOrdersCarpetRequest {
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

func (s *ServiceAdmin) SerachWithParam(req *pb.SerachWithParamRequest) error {
	if len(req.Mid) < consts.MinSearchMid && len(req.Url) < consts.MinSearchUrl {
		return errors.New(500, "REQUEST_INVALID", "Введите более точный параметр поиска")
	}

	return nil
}
