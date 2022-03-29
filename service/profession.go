package service

import (
	"context"
	"position_server/genproto/position_service"
	"position_server/package/helper"
	"position_server/package/logger"
	"position_server/storage"

	"github.com/jmoiron/sqlx"

	"google.golang.org/grpc/codes"
)

type professionService struct {
	logger  logger.Logger
	storage storage.StorageI
}

func NewProfessionService(log logger.Logger, db *sqlx.DB) *professionService {
	return &professionService{
		logger:  log,
		storage: storage.NewStoragePG(db),
	}
}

func (s *professionService) Create(ctx context.Context, req *position_service.CreateProfession) (*position_service.ProfessionId, error) {
	id, err := s.storage.Profession().Create(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while creating profession ", req, codes.Internal)
	}

	return &position_service.ProfessionId{
		Id: id,
	}, nil
}

func (s *professionService) Get(ctx context.Context, req *position_service.ProfessionId) (*position_service.Profession, error) {
	profession, err := s.storage.Profession().Get(req.Id)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting profession ", req, codes.Internal)
	}
	return profession, nil
}

func (s *professionService) GetAll(ctx context.Context, req *position_service.GetAllProfessionRequest) (*position_service.GetAllProfessionResponse, error) {
	professions, err := s.storage.Profession().GetAll(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting all profession ", req, codes.Internal)
	}

	return professions, nil
}

func (s *professionService) Update(ctx context.Context, req *position_service.Profession) (*position_service.Profession, error) {
	profession, err := s.storage.Profession().Update(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while update getting profession ", req, codes.Internal)
	}
	return profession, nil
}

func (s *professionService) Delete(ctx context.Context, req *position_service.ProfessionId) (*position_service.DeleteRes, error) {
	profession, err := s.storage.Profession().Delete(req.Id)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while delete getting profession ", req, codes.Internal)
	}
	return profession, nil
}
