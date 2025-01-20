package auth

import (
	"errors"
	"service_api/src/dtos"
	"service_api/src/entitys"
)

func (s *Service) Is(session string, ip string) bool {
	return s.cache.IsAdminCache(session, ip)
}

func (s *Service) Set(login string, password string, email string, ip string) (string, error) {
	admin := s.admin.GetAdmin(login, password, email)
	if admin == nil {
		return "", errors.New("Incorrect Admin param")
	}

	adminTransfer := &dtos.AdminTransfer{
		Id:    admin.ID,
		Login: admin.Login,
		Email: admin.Email,
		Role:  admin.Role,
		Ip:    ip,
	}

	cached := s.cache.SetAdminCache(adminTransfer)

	if len(cached) == 0 {
		return "", errors.New("Incorrect Append cache Admin param")
	}

	return cached, nil
}

func (s *Service) Get(session string) (*dtos.AdminTransfer, error) {
	adminTransfer, err := s.cache.GetAdminCache(session)
	if err != nil {
		return nil, errors.New("Session not found")
	}

	return adminTransfer, nil
}

func (s *Service) GetFull(session string) (*entitys.Admin, error) {
	adminTransfer, err := s.cache.GetAdminCache(session)
	if err != nil {
		return nil, errors.New("Session not found")
	}

	admin := s.admin.FindAdmin(adminTransfer.Id)
	if admin == nil {
		return nil, errors.New("Incorrect Admin Information")
	}

	admin.Pass = "*"

	return admin, nil
}

func (s *Service) FlushAll() {
	go s.cache.FlushAll()
}
