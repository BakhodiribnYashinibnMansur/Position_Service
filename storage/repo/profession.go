package repo

import "position_server/genproto/position_service"

type ProfessionRepoI interface {
	Create(req *position_service.CreateProfession) (string, error)
	Get(id string) (*position_service.Profession, error)
	GetAll(req *position_service.GetAllProfessionRequest) (*position_service.GetAllProfessionResponse, error)
	Update(req *position_service.Profession) (*position_service.Profession, error)
	Delete(id string) (*position_service.DeleteRes, error)
}
