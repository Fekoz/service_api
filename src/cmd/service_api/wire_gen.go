// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"bitbucket.org/klaraeng/package_config/conf"
	"bitbucket.org/klaraeng/package_queue/ampq"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"service_api/internal/auth"
	"service_api/internal/cache"
	"service_api/internal/components"
	"service_api/internal/components_admin"
	"service_api/internal/controller"
	"service_api/internal/persist"
	"service_api/internal/server"
	"service_api/internal/stop_list"
	"service_api/internal/validator"
)

// Injectors from wire.go:

// wireApp init application.
func wireApp(confServer *conf.Server, data *conf.Data, params *conf.Params, logger log.Logger) (*kratos.App, func(), error) {
	persistPersist, err := persist.NewPersist(data, logger)
	if err != nil {
		return nil, nil, err
	}
	masterEntityContext := persist.NewMasterEntity(persistPersist, logger)
	service := cache.NewCacheService(data, params)
	productContext := persist.NewProductRepo(persistPersist, logger)
	serviceCarpet := components.NewCarpetService(params, masterEntityContext, logger, service, productContext)
	showcasesContext := persist.NewShowcasesRepo(persistPersist, logger)
	serviceShowcases := components.NewShowcasesService(params, masterEntityContext, logger, showcasesContext)
	orderContext := persist.NewOrderRepo(persistPersist, logger)
	clientContext := persist.NewClientRepo(persistPersist, logger)
	queue := ampq.NewQueue(data, logger)
	mailTemplateContext := persist.NewMailTemplateRepo(persistPersist, logger)
	serviceOrderSelection := components.NewOrderSelectionService(params, orderContext, clientContext, logger, queue, mailTemplateContext)
	imagesContext := persist.NewImagesRepo(persistPersist, logger)
	priceContext := persist.NewPriceRepo(persistPersist, logger)
	serviceCatalog := components.NewCatalogService(params, masterEntityContext, imagesContext, priceContext, service, logger)
	serviceCarpets := components.NewCarpetsService(params, masterEntityContext, productContext, imagesContext, priceContext, logger)
	specificationContext := persist.NewSpecificationRepo(persistPersist, logger)
	attributeContext := persist.NewAttributeRepo(persistPersist, logger)
	collectionContext := persist.NewCollectionRepo(persistPersist, logger)
	serviceCarpetOnly := components.NewCarpetOnlyService(params, logger, productContext, imagesContext, priceContext, specificationContext, attributeContext, collectionContext, masterEntityContext)
	stopListContext := persist.NewStopListRepo(persistPersist, logger)
	stop_listService := stop_list.NewStopListService(stopListContext, params, service)
	serviceClient := validator.NewValidatorClientService()
	orderToCarpetContext := persist.NewOrderToCarpetRepo(persistPersist, logger)
	serviceOrderToCarpet := components.NewOrderToCarpetService(params, orderToCarpetContext, clientContext, productContext, priceContext, queue, mailTemplateContext, imagesContext, logger)
	categoryContext := persist.NewCategoryRepo(persistPersist, logger)
	serviceCategory := components.NewCategoryService(params, logger, service, categoryContext)
	viewcasesContext := persist.NewViewcasesRepo(persistPersist, logger)
	serviceViewcases := components.NewViewcasesService(params, masterEntityContext, service, logger, viewcasesContext)
	uriContext := persist.NewUriRepo(persistPersist, logger)
	serviceUtils := components.NewUtilsService(params, logger, uriContext)
	clientService := controller.NewClientService(serviceCarpet, serviceShowcases, serviceOrderSelection, serviceCatalog, serviceCarpets, serviceCarpetOnly, stop_listService, serviceClient, serviceOrderToCarpet, serviceCategory, serviceViewcases, serviceUtils)
	serviceAdmin := validator.NewValidatorAdminService()
	adminContext := persist.NewAdminRepo(persistPersist, logger)
	authService := auth.NewAuthService(adminContext, service, params)
	adminServiceAuth := components_admin.NewAdminAuthService(params, logger, authService, stop_listService)
	masterEntityAdminContext := persist.NewMasterEntityAdmin(persistPersist, logger)
	adminServiceCatalog := components_admin.NewAdminCatalogService(params, masterEntityAdminContext, imagesContext, priceContext, logger)
	adminServiceCarpet := components_admin.NewAdminCarpetService(params, masterEntityAdminContext, logger, productContext, priceContext, specificationContext, imagesContext, attributeContext)
	adminServiceCarpetOnly := components_admin.NewAdminCarpetOnlyService(params, logger, productContext, imagesContext, priceContext, masterEntityContext)
	adminServiceShowcases := components_admin.NewAdminShowcasesService(params, masterEntityAdminContext, logger, showcasesContext)
	adminServiceMailTemplate := components_admin.NewAdminMailTemplateService(params, logger, mailTemplateContext)
	marketSequenceContext := persist.NewMarketSequenceRepo(persistPersist, logger)
	adminServiceMarketSequence := components_admin.NewAdminMarketSequenceService(params, logger, priceContext, marketSequenceContext)
	adminServiceClient := components_admin.NewAdminClientService(params, logger, clientContext)
	adminServiceAdmin := components_admin.NewAdminAdminService(params, logger, adminContext)
	adminServiceOrder := components_admin.NewAdminOrderService(params, logger, orderContext)
	adminServiceOrderCarpet := components_admin.NewAdminOrderCarpetService(params, logger, orderToCarpetContext)
	adminServiceStopList := components_admin.NewAdminStopListService(params, logger, stopListContext)
	optionsContext := persist.NewOptionsRepo(persistPersist, logger)
	adminServiceOptions := components_admin.NewAdminOptionsService(params, logger, optionsContext)
	adminServiceViewcases := components_admin.NewAdminViewcasesService(params, masterEntityAdminContext, logger, viewcasesContext)
	reportsContext := persist.NewReportsRepo(persistPersist, logger)
	adminServiceReports := components_admin.NewAdminReportsService(params, reportsContext, logger)
	priceExportDynamicContext := persist.NewPriceExportDynamicRepo(persistPersist, logger)
	adminServicePriceDynamicExport := components_admin.NewAdminPriceDynamicExportService(params, logger, priceExportDynamicContext)
	adminService := controller.NewAdminService(serviceAdmin, adminServiceAuth, adminServiceCatalog, adminServiceCarpet, adminServiceCarpetOnly, adminServiceShowcases, adminServiceMailTemplate, adminServiceMarketSequence, adminServiceClient, adminServiceAdmin, adminServiceOrder, adminServiceOrderCarpet, adminServiceStopList, adminServiceOptions, adminServiceViewcases, adminServiceReports, adminServicePriceDynamicExport)
	grpcServer := server.NewGRPCServer(confServer, clientService, adminService, logger)
	app := newApp(logger, grpcServer)
	return app, func() {
	}, nil
}
