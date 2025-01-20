package controller

import (
	"context"
	"service_api/internal/components"
	"service_api/internal/stop_list"
	"service_api/internal/validator"
	"service_api/src/util/consts"

	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
)

type ClientService struct {
	pb.UnimplementedClientServer
	carpet        *components.ServiceCarpet
	showcase      *components.ServiceShowcases
	order         *components.ServiceOrderSelection
	catalog       *components.ServiceCatalog
	carpets       *components.ServiceCarpets
	carpetOnly    *components.ServiceCarpetOnly
	stopList      *stop_list.Service
	validator     *validator.ServiceClient
	orderToCarpet *components.ServiceOrderToCarpet
	category      *components.ServiceCategory
	viewcase      *components.ServiceViewcases
	utils         *components.ServiceUtils
}

func NewClientService(
	carpet *components.ServiceCarpet,
	showcase *components.ServiceShowcases,
	order *components.ServiceOrderSelection,
	catalog *components.ServiceCatalog,
	carpets *components.ServiceCarpets,
	carpetOnly *components.ServiceCarpetOnly,
	stopList *stop_list.Service,
	validator *validator.ServiceClient,
	orderToCarpet *components.ServiceOrderToCarpet,
	category *components.ServiceCategory,
	viewcase *components.ServiceViewcases,
	utils *components.ServiceUtils,
) *ClientService {
	return &ClientService{
		carpet:        carpet,
		showcase:      showcase,
		order:         order,
		catalog:       catalog,
		carpets:       carpets,
		carpetOnly:    carpetOnly,
		stopList:      stopList,
		validator:     validator,
		orderToCarpet: orderToCarpet,
		category:      category,
		viewcase:      viewcase,
		utils:         utils,
	}
}

func (s *ClientService) GetShowcase(ctx context.Context, req *pb.GetShowcaseRequest) (*pb.GetShowcaseReply, error) {
	_, err := s.validator.GetShowcase(req)
	if err != nil {
		return nil, err
	}
	reply, error := s.showcase.GetShowcase(req)
	if error != nil {
		return nil, error
	}
	return reply, nil
}

func (s *ClientService) GetCarpet(ctx context.Context, req *pb.GetCarpetRequest) (*pb.GetCarpetReply, error) {
	_, err := s.validator.GetCarpet(req)
	if err != nil {
		return nil, err
	}
	reply, error := s.carpet.GetCarpet(req)
	if error != nil {
		return nil, error
	}
	return reply, nil
}

func (s *ClientService) GetCarpetOnlyProduct(ctx context.Context, req *pb.GetCarpetOnlyProductRequest) (*pb.GetCarpetOnlyProductReply, error) {
	reply, error := s.carpetOnly.GetCarpetOnlyProduct(req)
	if error != nil {
		return nil, error
	}
	return reply, nil
}

func (s *ClientService) GetCarpetOnlyPrice(ctx context.Context, req *pb.GetCarpetOnlyPriceRequest) (*pb.GetCarpetOnlyPriceReply, error) {
	reply, error := s.carpetOnly.GetCarpetOnlyPrice(req)
	if error != nil {
		return nil, error
	}
	return reply, nil
}

func (s *ClientService) GetCarpetOnlySpecification(ctx context.Context, req *pb.GetCarpetOnlySpecificationRequest) (*pb.GetCarpetOnlySpecificationReply, error) {
	reply, error := s.carpetOnly.GetCarpetOnlySpecification(req)
	if error != nil {
		return nil, error
	}
	return reply, nil
}

func (s *ClientService) GetCarpetOnlyImages(ctx context.Context, req *pb.GetCarpetOnlyImagesRequest) (*pb.GetCarpetOnlyImagesReply, error) {
	reply, error := s.carpetOnly.GetCarpetOnlyImages(req)
	if error != nil {
		return nil, error
	}
	return reply, nil
}

func (s *ClientService) GetCarpetOnlyAttribute(ctx context.Context, req *pb.GetCarpetOnlyAttributeRequest) (*pb.GetCarpetOnlyAttributeReply, error) {
	reply, error := s.carpetOnly.GetCarpetOnlyAttribute(req)
	if error != nil {
		return nil, error
	}
	return reply, nil
}

func (s *ClientService) GetCarpets(ctx context.Context, req *pb.GetCarpetsRequest) (*pb.GetCarpetsReply, error) {
	req = s.validator.GetCarpets(req)
	reply, error := s.carpets.GetCarpets(req)
	if error != nil {
		return nil, error
	}
	return reply, nil
}

func (s *ClientService) CreateOrderLanding(ctx context.Context, req *pb.CreateOrderLandingRequest) (*pb.CreateOrderLandingReply, error) {
	_, err := s.validator.CreateOrderLanding(req)
	if err != nil {
		return nil, err
	}

	stopListErr := s.stopList.IsLock(ctx, consts.CreateOrderSelection)
	if stopListErr != nil {
		return nil, stopListErr
	}

	reply, error := s.order.Create(req)
	if error != nil {
		return nil, error
	}

	s.stopList.AppendLock(ctx, consts.CreateOrderSelection)
	return reply, nil
}

func (s *ClientService) CreateOrderToCarpet(ctx context.Context, req *pb.CreateOrderToCarpetRequest) (*pb.CreateOrderToCarpetReply, error) {
	_, err := s.validator.CreateOrderToCarpet(req)
	if err != nil {
		return nil, err
	}

	stopListErr := s.stopList.IsLock(ctx, consts.CreateOrderToCarpet)
	if stopListErr != nil {
		return nil, stopListErr
	}

	reply, error := s.orderToCarpet.Create(req)
	if error != nil {
		return nil, error
	}

	s.stopList.AppendLock(ctx, consts.CreateOrderToCarpet)
	return reply, nil
}

func (s *ClientService) CreateOrderToCarpetPrice(ctx context.Context, req *pb.CreateOrderToCarpetPriceRequest) (*pb.CreateOrderToCarpetPriceReply, error) {
	_, err := s.validator.CreateOrderToCarpetPrice(req)
	if err != nil {
		return nil, err
	}

	stopListErr := s.stopList.IsLock(ctx, consts.CreateOrderToCarpetPrice)
	if stopListErr != nil {
		return nil, stopListErr
	}

	reply, error := s.orderToCarpet.CreatePrice(req)
	if error != nil {
		return nil, error
	}

	s.stopList.AppendLock(ctx, consts.CreateOrderToCarpetPrice)
	return reply, nil
}

func (s *ClientService) GetCatalog(ctx context.Context, req *pb.GetCatalogRequest) (*pb.GetCatalogReply, error) {
	req = s.validator.GetCatalog(req)
	reply, error := s.catalog.GetCatalog(req)
	if error != nil {
		return nil, error
	}
	return reply, nil
}

func (s *ClientService) GetCatalogCategory(ctx context.Context, req *pb.GetCatalogCategoryRequest) (*pb.GetCatalogCategoryReply, error) {
	reply, error := s.category.GetCategory(req)
	if error != nil {
		return nil, error
	}
	return reply, nil
}

func (s *ClientService) GetCarpetCollection(ctx context.Context, req *pb.GetCarpetCollectionRequest) (*pb.GetCarpetCollectionReply, error) {
	reply, error := s.carpetOnly.GetCarpetCollection(req)
	if error != nil {
		return nil, error
	}
	return reply, nil
}

func (s *ClientService) GetViewcase(ctx context.Context, req *pb.GetViewcaseRequest) (*pb.GetViewcaseReply, error) {
	_, err := s.validator.GetViewcase(req)
	if err != nil {
		return nil, err
	}
	reply, error := s.viewcase.GetViewcase(req)
	if error != nil {
		return nil, error
	}
	return reply, nil
}

func (s *ClientService) GetMarketToUid(ctx context.Context, req *pb.GetMarketToUidRequest) (*pb.GetMarketToUidReply, error) {
	return s.utils.GetMarketToUid(req)
}
