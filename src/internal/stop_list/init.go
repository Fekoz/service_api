package stop_list

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewStopListService)
