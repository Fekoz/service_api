package components_admin

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewAdminShowcasesService,
	NewAdminAuthService,
	NewAdminCatalogService,
	NewAdminCarpetService,
	NewAdminCarpetOnlyService,
	NewAdminMailTemplateService,
	NewAdminMarketSequenceService,
	NewAdminClientService,
	NewAdminAdminService,
	NewAdminStopListService,
	NewAdminOrderService,
	NewAdminOrderCarpetService,
	NewAdminOptionsService,
	NewAdminViewcasesService,
	NewAdminReportsService,
	NewAdminPriceDynamicExportService,
)
