package components_admin

import (
	"encoding/json"
	"service_api/internal/persist"
	"service_api/src/entitys"

	conf "bitbucket.org/klaraeng/package_config/conf"

	pb "bitbucket.org/klaraeng/package_proto/service_api/admin"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type AdminServiceShowcases struct {
	pc           *conf.Params
	showcases    persist.ShowcasesContext
	masterEntity persist.MasterEntityAdminContext
	log          *log.Helper
}

func NewAdminShowcasesService(
	pc *conf.Params,
	masterEntity persist.MasterEntityAdminContext,
	logger log.Logger,
	showcases persist.ShowcasesContext,
) *AdminServiceShowcases {
	return &AdminServiceShowcases{
		pc:           pc,
		showcases:    showcases,
		masterEntity: masterEntity,
		log:          log.NewHelper(logger),
	}
}

func (s *AdminServiceShowcases) GetShowcasesAdmin(req *pb.GetShowcasesRequest) (*pb.GetShowcasesReply, error) {
	var showcaseList []*pb.GetShowcasesReply_Showcases
	count := s.showcases.FindListCount()

	if count < 1 {
		return nil, errors.New(500, "SHOWCASE_COUNT", "Пусто")
	}

	offset := req.Point*req.Limit - req.Limit
	array := s.showcases.FindList(req.Limit, offset)
	if array == nil {
		return nil, errors.New(500, "SHOWCASE_LIST", "Ошибка в получении списка витрин")
	}

	for _, el := range array {
		if el == nil {
			continue
		}

		var arr []string
		err := json.Unmarshal([]byte(el.PriceList), &arr)
		if err != nil {
			continue
		}

		showcaseList = append(showcaseList, &pb.GetShowcasesReply_Showcases{
			Name:    el.Name,
			Uuid:    el.Uuid,
			Count:   int64(len(arr)),
			Article: arr,
		})
	}

	return &pb.GetShowcasesReply{
		Success:  true,
		Limit:    req.Limit,
		Point:    req.Point,
		Count:    count,
		Showcase: showcaseList,
	}, nil
}

func (s *AdminServiceShowcases) GetShowcase(req *pb.GetShowcaseRequest) (*pb.GetShowcaseReply, error) {
	var mel []*pb.MasterEntityLarge
	r := s.showcases.FindByUUIDAll(req.Uuid)
	if r == nil {
		return nil, errors.New(500, "UUID_CLEARED", "Не найдено значение с переданным UUID")
	}
	if len(r.PriceList) < 1 {
		return nil, errors.New(500, "PRICE_LIST_CLEARED", "Пустой список, передаваемый в Price List")
	}
	var arr []string
	err := json.Unmarshal([]byte(r.PriceList), &arr)
	if err != nil {
		return nil, errors.New(500, "PRICE_LIST_CLEARED", "Не корреткно указано значение Price List")
	}
	for _, el := range arr {
		if len(el) < 1 {
			continue
		}
		me := s.masterEntity.OneToPriceIdLarge(el)
		if me != nil {
			mel = append(mel, me)
		}
	}
	return &pb.GetShowcaseReply{
		Success:      true,
		Name:         r.Name,
		MasterEntity: mel,
		IsActive:     r.IsActive,
	}, nil
}

func (s *AdminServiceShowcases) SetShowcase(req *pb.SetShowcaseRequest) (*pb.SetShowcaseReply, error) {
	if len(req.Uuid) < 1 {
		return nil, errors.New(500, "REQUEST_HAS_NULL", "Не указаны переданные значения UUID")
	}

	showcase := s.showcases.FindByUUIDAll(req.Uuid)
	array, err := json.Marshal(req.PriceId)
	if err != nil {
		return nil, errors.New(500, "REQUEST_HAS_NULL", "Некорректно передан PRICE_LIST")
	}

	if showcase == nil || showcase.ID == 0 {
		_ = s.showcases.Create(&entitys.Showcases{
			Name:      req.Name,
			IsActive:  req.IsActive,
			PriceList: string(array),
			Uuid:      uuid.NewString(),
		})
	} else {
		showcase.IsActive = req.IsActive
		showcase.Name = req.Name
		showcase.PriceList = string(array)
		_ = s.showcases.Update(showcase)
	}

	return &pb.SetShowcaseReply{
		Success: true,
	}, nil
}

func (s *AdminServiceShowcases) DropShowcase(req *pb.DropShowcaseRequest) (*pb.DropShowcaseReply, error) {
	if len(req.Uuid) < 1 {
		return nil, errors.New(500, "REQUEST_HAS_NULL", "Не указаны переданные значения UUID")
	}

	return &pb.DropShowcaseReply{
		Success: s.showcases.Drop(req.Uuid),
	}, nil
}
