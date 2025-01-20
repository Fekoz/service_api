package components_admin

import (
	"encoding/json"
	"service_api/internal/persist"
	"strconv"

	conf "bitbucket.org/klaraeng/package_config/conf"

	pb "bitbucket.org/klaraeng/package_proto/service_api/admin"

	"github.com/go-kratos/kratos/v2/log"
)

type AdminServiceReports struct {
	pc  *conf.Params
	rpt persist.ReportsContext
	log *log.Helper
}

func NewAdminReportsService(
	pc *conf.Params,
	rpt persist.ReportsContext,
	logger log.Logger,
) *AdminServiceReports {
	return &AdminServiceReports{
		pc:  pc,
		rpt: rpt,
		log: log.NewHelper(logger),
	}
}

func (s *AdminServiceReports) reportMarketPlace(limit int64, point int64, option []*pb.GetReportRequest_Option) *pb.GetReportReply {
	var minCount int64 = 0
	for _, el := range option {
		if el.Key == "min-count" && len(el.Value) > 0 {
			minCount, _ = strconv.ParseInt(el.Value, 10, 64)
		}
	}

	count := s.rpt.FindByMarketPlaceCount(minCount, limit, point)
	data := s.rpt.FindByMarketPlace(minCount, limit, point*limit)

	isSuccess := false
	if len(data) > 0 && data != nil {
		isSuccess = true
	}

	marshal, err := json.Marshal(data)
	if err != nil {
		marshal = []byte("{}")
	}

	return &pb.GetReportReply{
		Success: isSuccess,
		Option: &pb.GetReportReply_Option{
			Limit: limit,
			Point: point,
			Count: count / limit,
			Data:  string(marshal),
		},
	}
}

func (s *AdminServiceReports) GetReport(req *pb.GetReportRequest) *pb.GetReportReply {
	switch req.Type {
	case "marketplace":
		return s.reportMarketPlace(req.Limit, req.Point, req.Options)
	}
	return &pb.GetReportReply{}
}
