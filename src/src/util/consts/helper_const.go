package consts

const (
	MetadataAdminAuth string = "x-md-global-admin-auth"
	MaskCacheKey      string = "%s_%s"
	MetadataIP        string = "x-md-global-client-ip"
	MetadataIPProxy   string = "x-md-global-client-proxy-ip"
	// FILTER VALIDATOR
	MaxLimit     int64 = 40
	MinLimit     int64 = 10
	DefaultPoint int64 = 1
	MinSearchMid int   = 14
	MinSearchUrl int   = 36
	// STOP LIST SERVICE
	MaxTry                   int    = 5
	DayPer                   int    = 1
	CreateOrderSelection     string = "CreateOrderSelection.Service"
	CreateOrderToCarpet      string = "CreateOrderToCarpet.Service"
	CreateOrderToCarpetPrice string = "CreateOrderToCarpetPrice.Service"
	CreatAdminAuth           string = "AdminAuth.Service"

	PacketSizeInsertOrderField int = 6

	// STATUS TO ORDER
	OrderStatusClose         int64 = 0
	OrderStatusStart         int64 = 1
	OrderStatusApprove       int64 = 2
	OrderStatusPay           int64 = 3
	OrderStatusReserve       int64 = 4
	OrderStatusWait          int64 = 5
	OrderStatusSend          int64 = 6
	OrderStatusFinal         int64 = 7
	OrderStatusDisputeOpen   int64 = -1
	OrderStatusDisputeReturn int64 = -2
	OrderStatusDisputePay    int64 = -3

	// TYPE TO ORDER
	OrderTypeSelection      int64 = 1
	OrderTypeCommission     int64 = 2
	OrderTypeCommissionFull int64 = 3

	// QUEUE NAMED
	QueueOther    int64 = 0
	QueueEmail    int64 = 1
	QueueTelegram int64 = 2

	// CARPET CONSTS
	CARPET_FORM        string = "FORM"
	CARPET_DESIGN      string = "DESIGN"
	CARPET_COLOR       string = "COLOR"
	CARPET_PAY_LIMIT   string = "PAY_LIMIT"
	CARPET_SIZE_WIDTH  string = "SIZE_WIDTH"
	CARPET_SIZE_HEIGHT string = "SIZE_HEIGHT"

	CARPET_PRESENT_ORDER_SELECTION string = "SALE10"

	DefaultImageDir  string = "/export/"
	DefaultImageFile string = "defaultImage.jpg"

	ParserVeneraEmail    string = "parser.venera.email"
	ParserVeneraPassword string = "parser.venera.password"
	ParserDefaultFactor  string = "parser.default.factor"
	PriceFactor          string = "price.factor"
	PriceMarkup          string = "price.markup"
	PriceRandom          string = "price.random"
	PriceRandomMin       string = "price.random.min"
	PriceRandomMax       string = "price.random.max"
	PriceOldFactor       string = "price.old.factor"
	PriceOldMarkup       string = "price.old.markup"
	PriceOldRandom       string = "price.old.random"
	PriceOldRandomMin    string = "price.old.random.min"
	PriceOldRandomMax    string = "price.old.random.max"

	//ARRAY TYPES CATGORY
	TypeCastValue int = 0
	TypeCastKey   int = 1
	TypeCastType  int = 2

	// QUEUE CONST
	QUEUE_DURABLE     bool = true
	QUEUE_PASSIVE     bool = false
	QUEUE_EXCLUSIVE   bool = false
	QUEUE_AUTO_DELETE bool = false
	QUEUE_NO_LOCAL    bool = false
	QUEUE_NO_ACK      bool = false
	QUEUE_NO_WAIT     bool = false
)

var StatusOrder = func() map[int64]string {
	return map[int64]string{
		OrderStatusClose:         "Заказ закрыт",
		OrderStatusStart:         "Заказ зарегистрирован",
		OrderStatusApprove:       "Заказ на подтверждении",
		OrderStatusPay:           "Заказ оплачен",
		OrderStatusReserve:       "Заказ в резервации товара",
		OrderStatusWait:          "Заказ в ожидании отправки",
		OrderStatusSend:          "Заказ отправлен",
		OrderStatusFinal:         "Заказ доставлен",
		OrderStatusDisputeOpen:   "Открыт спор по заказу",
		OrderStatusDisputeReturn: "Заказ забрали у клиента",
		OrderStatusDisputePay:    "Выданы деньги клиенту",
	}
}

var TypeOrder = func() map[int64]string {
	return map[int64]string{
		OrderTypeSelection:      "Подбор заказа",
		OrderTypeCommission:     "Заказ ковра",
		OrderTypeCommissionFull: "Индивидуальный заказ ковра с размерами",
	}
}

var QueueList = func() map[int64]string {
	return map[int64]string{
		QueueOther:    "other",
		QueueEmail:    "email",
		QueueTelegram: "telegram",
	}
}
