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

type AdminServiceViewcases struct {
	pc           *conf.Params
	viewcases    persist.ViewcasesContext
	masterEntity persist.MasterEntityAdminContext
	log          *log.Helper
}

func NewAdminViewcasesService(
	pc *conf.Params,
	masterEntity persist.MasterEntityAdminContext,
	logger log.Logger,
	viewcases persist.ViewcasesContext,
) *AdminServiceViewcases {
	return &AdminServiceViewcases{
		pc:           pc,
		viewcases:    viewcases,
		masterEntity: masterEntity,
		log:          log.NewHelper(logger),
	}
}

func (s *AdminServiceViewcases) GetViewcases(req *pb.GetViewcasesRequest) (*pb.GetViewcasesReply, error) {
	var ViewcaseList []*pb.GetViewcasesReply_Viewcases
	count := s.viewcases.FindListCount()

	if count < 1 {
		return nil, errors.New(500, "Viewcase_COUNT", "Пусто")
	}

	offset := req.Point*req.Limit - req.Limit
	array := s.viewcases.FindList(req.Limit, offset)
	if array == nil {
		return nil, errors.New(500, "Viewcase_LIST", "Ошибка в получении списка витрин")
	}

	for _, el := range array {
		if el == nil {
			continue
		}

		var arr []string
		err := json.Unmarshal([]byte(el.List), &arr)
		if err != nil {
			continue
		}

		ViewcaseList = append(ViewcaseList, &pb.GetViewcasesReply_Viewcases{
			Name:    el.Name,
			Uuid:    el.Uuid,
			Count:   int64(len(arr)),
			Article: arr,
		})
	}

	return &pb.GetViewcasesReply{
		Success:  true,
		Limit:    req.Limit,
		Point:    req.Point,
		Count:    count,
		Viewcase: ViewcaseList,
	}, nil
}

func (s *AdminServiceViewcases) GetViewcase(req *pb.GetViewcaseRequest) (*pb.GetViewcaseReply, error) {
	var mel []*pb.MasterEntity
	r := s.viewcases.FindByUUIDAll(req.Uuid)
	if r == nil {
		return nil, errors.New(500, "UUID_CLEARED", "Не найдено значение с переданным UUID")
	}
	if len(r.List) < 1 {
		return nil, errors.New(500, "PRICE_LIST_CLEARED", "Пустой список, передаваемый в List")
	}
	var arr []string
	err := json.Unmarshal([]byte(r.List), &arr)
	if err != nil {
		return nil, errors.New(500, "PRICE_LIST_CLEARED", "Не корреткно указано значение List")
	}
	for _, el := range arr {
		if len(el) < 1 {
			continue
		}
		me := s.masterEntity.OneToId(s.masterEntity.ArticleToProductId(el))
		if me != nil &&
			me.Product != nil &&
			me.Price != nil &&
			me.Specification != nil && len(me.Specification) > 0 &&
			me.Images != nil {
			mel = append(mel, me)
		}
	}
	return &pb.GetViewcaseReply{
		Success:      true,
		Name:         r.Name,
		MasterEntity: mel,
		IsActive:     r.IsActive,
	}, nil
}

func (s *AdminServiceViewcases) SetViewcase(req *pb.SetViewcaseRequest) (*pb.SetViewcaseReply, error) {
	if len(req.Uuid) < 1 {
		return nil, errors.New(500, "REQUEST_HAS_NULL", "Не указаны переданные значения UUID")
	}

	Viewcase := s.viewcases.FindByUUIDAll(req.Uuid)
	array, err := json.Marshal(req.Articles)
	if err != nil {
		return nil, errors.New(500, "REQUEST_HAS_NULL", "Некорректно передан LIST")
	}

	if Viewcase == nil || Viewcase.ID == 0 {
		_ = s.viewcases.Create(&entitys.Viewcases{
			Name:     req.Name,
			IsActive: req.IsActive,
			List:     string(array),
			Uuid:     uuid.NewString(),
		})
	} else {
		Viewcase.IsActive = req.IsActive
		Viewcase.Name = req.Name
		Viewcase.List = string(array)
		_ = s.viewcases.Update(Viewcase)
	}

	return &pb.SetViewcaseReply{
		Success: true,
	}, nil
}

func (s *AdminServiceViewcases) DropViewcase(req *pb.DropViewcaseRequest) (*pb.DropViewcaseReply, error) {
	if len(req.Uuid) < 1 {
		return nil, errors.New(500, "REQUEST_HAS_NULL", "Не указаны переданные значения UUID")
	}

	return &pb.DropViewcaseReply{
		Success: s.viewcases.Drop(req.Uuid),
	}, nil
}
