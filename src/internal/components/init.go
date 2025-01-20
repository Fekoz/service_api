package components

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewCarpetService,
	NewShowcasesService,
	NewCatalogService,
	NewCarpetsService,
	NewCarpetOnlyService,
	NewOrderSelectionService,
	NewOrderToCarpetService,
	NewCategoryService,
	NewViewcasesService,
	NewUtilsService,
)
