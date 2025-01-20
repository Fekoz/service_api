package dtos

type MarketTypesTransferPackage struct {
	Limit int64
	Point int64
	Count int64
	Data  []MarketTypesTransfer
}

type MarketTypesTransfer struct {
	Article     string
	Width       string
	Height      string
	TotalCount  bool
	Mid         string
	TypeMarket  int
	Url         string
	Marketplace string
}
