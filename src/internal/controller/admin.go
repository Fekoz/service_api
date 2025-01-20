package controller

import (
	"context"
	"service_api/internal/components_admin"
	"service_api/internal/validator"

	pb "bitbucket.org/klaraeng/package_proto/service_api/admin"
)

type AdminService struct {
	pb.UnimplementedAdminServer
	validator      *validator.ServiceAdmin
	auth           *components_admin.AdminServiceAuth
	showcase       *components_admin.AdminServiceShowcases
	catalog        *components_admin.AdminServiceCatalog
	carpet         *components_admin.AdminServiceCarpet
	carpetOnly     *components_admin.AdminServiceCarpetOnly
	mailTemplate   *components_admin.AdminServiceMailTemplate
	marketSequence *components_admin.AdminServiceMarketSequence
	client         *components_admin.AdminServiceClient
	admin          *components_admin.AdminServiceAdmin
	order          *components_admin.AdminServiceOrder
	orderCarpet    *components_admin.AdminServiceOrderCarpet
	stopList       *components_admin.AdminServiceStopList
	options        *components_admin.AdminServiceOptions
	viewcase       *components_admin.AdminServiceViewcases
	reports        *components_admin.AdminServiceReports
	pdea           *components_admin.AdminServicePriceDynamicExport
}

func NewAdminService(
	validator *validator.ServiceAdmin,
	auth *components_admin.AdminServiceAuth,
	catalog *components_admin.AdminServiceCatalog,
	carpet *components_admin.AdminServiceCarpet,
	carpetOnly *components_admin.AdminServiceCarpetOnly,
	showcase *components_admin.AdminServiceShowcases,
	mailTemplate *components_admin.AdminServiceMailTemplate,
	marketSequence *components_admin.AdminServiceMarketSequence,
	client *components_admin.AdminServiceClient,
	admin *components_admin.AdminServiceAdmin,
	order *components_admin.AdminServiceOrder,
	orderCarpet *components_admin.AdminServiceOrderCarpet,
	stopList *components_admin.AdminServiceStopList,
	options *components_admin.AdminServiceOptions,
	viewcase *components_admin.AdminServiceViewcases,
	reports *components_admin.AdminServiceReports,
	pdea *components_admin.AdminServicePriceDynamicExport,
) *AdminService {
	return &AdminService{
		validator:      validator,
		auth:           auth,
		showcase:       showcase,
		catalog:        catalog,
		carpet:         carpet,
		carpetOnly:     carpetOnly,
		mailTemplate:   mailTemplate,
		marketSequence: marketSequence,
		client:         client,
		admin:          admin,
		order:          order,
		orderCarpet:    orderCarpet,
		stopList:       stopList,
		options:        options,
		viewcase:       viewcase,
		reports:        reports,
		pdea:           pdea,
	}
}

func (s *AdminService) Auth(ctx context.Context, req *pb.GetAuthRequest) (*pb.GetAuthReply, error) {
	return s.auth.Auth(ctx, req)
}

func (s *AdminService) IsAuth(ctx context.Context, req *pb.IsAuthRequest) (*pb.IsAuthReply, error) {
	isAuth := false
	if nil == s.auth.IsAuth(ctx) {
		isAuth = true
	}
	return &pb.IsAuthReply{Success: isAuth}, nil
}

func (s *AdminService) GetAdmin(ctx context.Context, req *pb.GetAdminRequest) (*pb.GetAdminReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	return s.auth.GetAdmin(ctx)
}

func (s *AdminService) GetAdminInformation(ctx context.Context, req *pb.GetAdminInformationRequest) (*pb.GetAdminInformationReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	return s.auth.GetAdminInformation(ctx)
}

func (s *AdminService) GetShowcases(ctx context.Context, req *pb.GetShowcasesRequest) (*pb.GetShowcasesReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	req = s.validator.GetShowcases(req)
	return s.showcase.GetShowcasesAdmin(req)
}

func (s *AdminService) GetShowcase(ctx context.Context, req *pb.GetShowcaseRequest) (*pb.GetShowcaseReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	return s.showcase.GetShowcase(req)
}

func (s *AdminService) SetShowcase(ctx context.Context, req *pb.SetShowcaseRequest) (*pb.SetShowcaseReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	return s.showcase.SetShowcase(req)
}

func (s *AdminService) DropShowcase(ctx context.Context, req *pb.DropShowcaseRequest) (*pb.DropShowcaseReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	return s.showcase.DropShowcase(req)
}

func (s *AdminService) GetCatalog(ctx context.Context, req *pb.GetCatalogRequest) (*pb.GetCatalogReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	req = s.validator.GetCatalog(req)
	reply, err := s.catalog.GetCatalog(req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *AdminService) GetCarpetOnlyPrice(ctx context.Context, req *pb.GetCarpetOnlyPriceRequest) (*pb.GetCarpetOnlyPriceReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	reply, err := s.carpetOnly.GetCarpetOnlyPrice(req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *AdminService) GetCarpetMinToPriceUUID(ctx context.Context, req *pb.GetCarpetMinToPriceUUIDRequest) (*pb.GetCarpetMinToPriceUUIDReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	reply, err := s.carpetOnly.GetCarpetMinToPriceUUID(req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *AdminService) GetCarpetMinToPriceMID(ctx context.Context, req *pb.GetCarpetMinToPriceMIDRequest) (*pb.GetCarpetMinToPriceMIDReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	reply, err := s.carpetOnly.GetCarpetMinToPriceMID(req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *AdminService) GetCarpetMinToArticle(ctx context.Context, req *pb.GetCarpetMinToArticleRequest) (*pb.GetCarpetMinToArticleReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	reply, err := s.carpetOnly.GetCarpetMinToArticle(req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *AdminService) GetCarpet(ctx context.Context, req *pb.GetCarpetRequest) (*pb.GetCarpetReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	reply, err := s.carpet.GetCarpet(req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *AdminService) SetCarpet(ctx context.Context, req *pb.SetCarpetRequest) (*pb.SetCarpetReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	reply, err := s.carpet.SetCarpet(req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *AdminService) GetMarketSequences(ctx context.Context, req *pb.GetMarketSequencesRequest) (*pb.GetMarketSequencesReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	req = s.validator.GetMarketSequences(req)
	return s.marketSequence.GetMarketSequences(req)
}

func (s *AdminService) SetMarketSequence(ctx context.Context, req *pb.SetMarketSequenceRequest) (*pb.SetMarketSequenceReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	err := s.validator.SetMarketSequence(req)
	if err != nil {
		return nil, err
	}
	return s.marketSequence.SetMarketSequence(req)
}

func (s *AdminService) GetMailTemplates(ctx context.Context, req *pb.GetMailTemplatesRequest) (*pb.GetMailTemplatesReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	req = s.validator.GetMailTemplates(req)
	return s.mailTemplate.GetMailTemplates(req)
}

func (s *AdminService) GetMailTemplate(ctx context.Context, req *pb.GetMailTemplateRequest) (*pb.GetMailTemplateReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	return s.mailTemplate.GetMailTemplate(req)
}

func (s *AdminService) SetMailTemplate(ctx context.Context, req *pb.SetMailTemplateRequest) (*pb.SetMailTemplateReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	err := s.validator.SetMailTemplate(req)
	if err != nil {
		return nil, err
	}
	return s.mailTemplate.SetMailTemplate(req)
}

func (s *AdminService) GetOrders(ctx context.Context, req *pb.GetOrdersRequest) (*pb.GetOrdersReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	req = s.validator.GetOrders(req)
	return s.order.GetOrders(req)
}

func (s *AdminService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	return s.order.GetOrder(req)
}

func (s *AdminService) SetOrder(ctx context.Context, req *pb.SetOrderRequest) (*pb.SetOrderReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	return s.order.SetOrder(req)
}

func (s *AdminService) GetOrdersCarpet(ctx context.Context, req *pb.GetOrdersCarpetRequest) (*pb.GetOrdersCarpetReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	req = s.validator.GetOrdersCarpet(req)
	return s.orderCarpet.GetOrdersCarpet(req)
}

func (s *AdminService) GetOrderCarpet(ctx context.Context, req *pb.GetOrderCarpetRequest) (*pb.GetOrderCarpetReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	return s.orderCarpet.GetOrderCarpet(req)
}

func (s *AdminService) SetOrderCarpet(ctx context.Context, req *pb.SetOrderCarpetRequest) (*pb.SetOrderCarpetReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	return s.orderCarpet.SetOrderCarpet(req)
}

func (s *AdminService) GetClients(ctx context.Context, req *pb.GetClientsRequest) (*pb.GetClientsReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	req = s.validator.GetClients(req)
	reply, err := s.client.GetClients(req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *AdminService) GetAdmins(ctx context.Context, req *pb.GetAdminsRequest) (*pb.GetAdminsReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	req = s.validator.GetAdmins(req)
	reply, err := s.admin.GetAdmins(req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *AdminService) SetAdmin(ctx context.Context, req *pb.SetAdminRequest) (*pb.SetAdminReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	err := s.validator.SetAdmin(req)
	if err != nil {
		return nil, err
	}
	return s.admin.SetAdmin(req)
}

func (s *AdminService) GetStopLists(ctx context.Context, req *pb.GetStopListsRequest) (*pb.GetStopListsReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	req = s.validator.GetStopLists(req)
	return s.stopList.GetStopLists(req)
}

func (s *AdminService) DropStopList(ctx context.Context, req *pb.DropStopListRequest) (*pb.DropStopListReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	return s.stopList.DropStopList(req)
}

func (s *AdminService) CacheClear(ctx context.Context, req *pb.CacheClear) (*pb.CacheClear, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	s.auth.CacheClear()
	return &pb.CacheClear{}, nil
}

func (s *AdminService) OptionsChange(ctx context.Context, req *pb.OptionsChangeRequest) (*pb.OptionsChangeReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	return s.options.Change(req)
}

func (s *AdminService) OptionsList(ctx context.Context, req *pb.OptionsListRequest) (*pb.OptionsListReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	return s.options.List()
}

func (s *AdminService) GetViewcases(ctx context.Context, req *pb.GetViewcasesRequest) (*pb.GetViewcasesReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	return s.viewcase.GetViewcases(req)
}

func (s *AdminService) GetViewcase(ctx context.Context, req *pb.GetViewcaseRequest) (*pb.GetViewcaseReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	return s.viewcase.GetViewcase(req)
}

func (s *AdminService) SetViewcase(ctx context.Context, req *pb.SetViewcaseRequest) (*pb.SetViewcaseReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	return s.viewcase.SetViewcase(req)
}

func (s *AdminService) DropViewcase(ctx context.Context, req *pb.DropViewcaseRequest) (*pb.DropViewcaseReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	return s.viewcase.DropViewcase(req)
}

func (s *AdminService) GetReport(ctx context.Context, req *pb.GetReportRequest) (*pb.GetReportReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	return s.reports.GetReport(req), nil
}

func (s *AdminService) GetPriceExportDynamics(ctx context.Context, req *pb.GetPriceExportDynamicsRequest) (*pb.GetPriceExportDynamicsReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	return s.pdea.List(req)
}

func (s *AdminService) SetPriceExportDynamic(ctx context.Context, req *pb.SetPriceExportDynamicRequest) (*pb.SetPriceExportDynamicReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	return s.pdea.Set(req)
}

func (s *AdminService) DropPriceExportDynamic(ctx context.Context, req *pb.DropPriceExportDynamicRequest) (*pb.DropPriceExportDynamicReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}
	return s.pdea.Drop(req)
}

func (s *AdminService) SerachWithParam(ctx context.Context, req *pb.SerachWithParamRequest) (*pb.SerachWithParamReply, error) {
	auth := s.auth.IsAuth(ctx)
	if auth != nil {
		return nil, auth
	}

	validator := s.validator.SerachWithParam(req)
	if validator != nil {
		return nil, validator
	}

	reply, err := s.catalog.SerachWithParam(req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
