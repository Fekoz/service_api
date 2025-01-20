package components_admin

import (
	"service_api/internal/persist"

	conf "bitbucket.org/klaraeng/package_config/conf"

	pb "bitbucket.org/klaraeng/package_proto/service_api/admin"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type AdminServiceMailTemplate struct {
	pc           *conf.Params
	mailTemplate persist.MailTemplateContext
	log          *log.Helper
}

func NewAdminMailTemplateService(
	pc *conf.Params,
	logger log.Logger,
	mailTemplate persist.MailTemplateContext) *AdminServiceMailTemplate {
	return &AdminServiceMailTemplate{
		pc:           pc,
		mailTemplate: mailTemplate,
		log:          log.NewHelper(logger),
	}
}

func (s *AdminServiceMailTemplate) GetMailTemplates(req *pb.GetMailTemplatesRequest) (*pb.GetMailTemplatesReply, error) {
	var array []*pb.GetMailTemplatesReply_MailTemplates
	var count int64
	count = s.mailTemplate.FindListCount()

	if count < 1 {
		return nil, errors.New(500, "MAIL_TEMPLATE", "Пусто")
	}

	offset := req.Point*req.Limit - req.Limit

	qb := s.mailTemplate.FindList(req.Limit, offset)
	if qb == nil {
		return nil, errors.New(500, "MAIL_TEMPLATE", "Ошибка в получении списка")
	}

	for _, el := range qb {
		if el == nil {
			continue
		}
		array = append(array, &pb.GetMailTemplatesReply_MailTemplates{
			Id:    int64(el.ID),
			Title: el.Title,
			Date:  el.UpdateAt.String(),
		})
	}

	return &pb.GetMailTemplatesReply{
		Success:       true,
		Limit:         req.Limit,
		Point:         req.Point,
		Count:         count,
		MailTemplates: array,
	}, nil
}

func (s *AdminServiceMailTemplate) GetMailTemplate(req *pb.GetMailTemplateRequest) (*pb.GetMailTemplateReply, error) {
	mailTemplate := s.mailTemplate.FindByID(req.Id)
	if mailTemplate == nil {
		return nil, errors.New(500, "MAIL_TEMPLATE", "Ошибка в получении шаблона")
	}

	return &pb.GetMailTemplateReply{
		Success: true,
		Id:      int64(mailTemplate.ID),
		Title:   mailTemplate.Title,
		Message: mailTemplate.Message,
		Date:    mailTemplate.UpdateAt.String(),
	}, nil
}

func (s *AdminServiceMailTemplate) SetMailTemplate(req *pb.SetMailTemplateRequest) (*pb.SetMailTemplateReply, error) {
	isUpdate := s.mailTemplate.SetByID(req.Id, req.Title, req.Message)
	if true == isUpdate {
		return &pb.SetMailTemplateReply{Success: true}, nil
	}

	return nil, errors.New(500, "MAIL_TEMPLATE", "Ошибка при обновлении")
}
