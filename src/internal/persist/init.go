package persist

import (
	conf "bitbucket.org/klaraeng/package_config/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	// ALL Repo
	NewPersist,
	NewProductRepo,
	NewPriceRepo,
	NewSpecificationRepo,
	NewAttributeRepo,
	NewImagesRepo,
	NewMasterEntity,
	NewShowcasesRepo,
	NewOrderRepo,
	NewOrderToCarpetRepo,
	NewStopListRepo,
	NewClientRepo,
	NewSenderRepo,
	NewCategoryRepo,
	NewAdminRepo,
	NewMarketSequenceRepo,
	NewMailTemplateRepo,
	NewCollectionRepo,
	NewViewcasesRepo,
	NewUriRepo,
	NewPriceExportDynamicRepo,
	// Only Admin Repo
	NewMasterEntityAdmin,
	NewOptionsRepo,
	// Reports
	NewReportsRepo,
)

type Persist struct {
	pdb *gorm.DB
	log *log.Helper
}

func NewPersist(c *conf.Data, logger log.Logger) (*Persist, error) {
	dsn :=
		"host=" + c.Database.Source +
			" user=" + c.Database.Login +
			" password=" + c.Database.Password +
			" dbname=" + c.Database.Database +
			" port=" + c.Database.Port +
			" sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &Persist{
		pdb: db,
		log: log.NewHelper(logger),
	}, nil
}
