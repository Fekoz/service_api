package components_admin

import (
	"fmt"
	"service_api/internal/persist"
	"service_api/src/entitys"
	"service_api/src/util"
	"service_api/src/util/consts"
	"strconv"

	conf "bitbucket.org/klaraeng/package_config/conf"

	pb "bitbucket.org/klaraeng/package_proto/service_api/admin"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type AdminServiceOptions struct {
	pc      *conf.Params
	options persist.OptionsContext
	log     *log.Helper
}

func NewAdminOptionsService(
	pc *conf.Params,
	logger log.Logger,
	options persist.OptionsContext) *AdminServiceOptions {
	return &AdminServiceOptions{
		pc:      pc,
		options: options,
		log:     log.NewHelper(logger),
	}
}

func (s *AdminServiceOptions) append(option *entitys.Options, value string, info string) {
	option.Value = value
	option.Info = info
	s.options.Update(option)
}

func (s *AdminServiceOptions) Change(req *pb.OptionsChangeRequest) (*pb.OptionsChangeReply, error) {
	var optionEm *entitys.Options
	optionEm = s.options.Find(req.Name)
	if optionEm == nil {
		return nil, errors.New(500, "OPTIONS", "Передан некорректный атрибут")
	}

	switch req.Name {
	case
		consts.ParserDefaultFactor,
		consts.PriceFactor,
		consts.PriceOldFactor:
		val, err := strconv.ParseFloat(req.Option.Value, 32)
		if err != nil {
			return nil, errors.New(500, "OPTIONS", fmt.Sprintf("Передан некорректный формат (не дробный) для атрибута [%s]", req.Name))
		}
		value := util.ToFixed(val, 2)
		s.append(optionEm, fmt.Sprintf("%.2f", value), req.Option.Info)

	case
		consts.PriceMarkup,
		consts.PriceRandomMin,
		consts.PriceRandomMax,
		consts.PriceOldMarkup,
		consts.PriceOldRandomMin,
		consts.PriceOldRandomMax:
		_, err := strconv.ParseInt(req.Option.Value, 10, 64)
		if err != nil {
			return nil, errors.New(500, "OPTIONS", fmt.Sprintf("Передан некорректный формат (не числовой) для атрибута [%s]", req.Name))
		}
		s.append(optionEm, req.Option.Value, req.Option.Info)

	default:
		s.append(optionEm, req.Option.Value, req.Option.Info)

	}

	return &pb.OptionsChangeReply{
		Success: true,
	}, nil
}

func (s *AdminServiceOptions) List() (*pb.OptionsListReply, error) {
	var options []*pb.Options

	for _, option := range s.options.List() {
		options = append(options, &pb.Options{
			Name:  option.Name,
			Value: option.Value,
			Info:  option.Info,
			Date:  option.UpdateAt.String(),
		})
	}

	var isSuccess bool
	isSuccess = false
	if len(options) > 1 {
		isSuccess = true
	}

	return &pb.OptionsListReply{
		Success: isSuccess,
		Option:  options,
	}, nil
}
