package controller

import (
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/logic_error"
	"fmt"
)

type PictureController struct {
	repo repository.IPictureRepo
	uc   controller.IUserController
}

func (p *PictureController) GetListByDorama(idD int) ([]model.Picture, error) {
	res, err := p.repo.GetListDorama(idD)
	if err != nil {
		return nil, fmt.Errorf("getByDorama: %w", err)
	}
	return res, nil
}

func (p *PictureController) GetListByStaff(idS int) ([]model.Picture, error) {
	res, err := p.repo.GetListStaff(idS)
	if err != nil {
		return nil, fmt.Errorf("getByStaff: %w", err)
	}
	return res, nil
}

func (p *PictureController) CreatePicture(token string, record model.Picture) error {
	user, err := p.uc.AuthByToken(token)
	if err != nil {
		return fmt.Errorf("authToken: %w", err)
	}

	//TODO +adminAccessError
	if !user.IsAdmin {
		return fmt.Errorf("%w", logic_error.ErrorAdminAccess)
	}

	err = p.repo.CreatePicture(record)
	if err != nil {
		return fmt.Errorf("createPicture: %w", err)
	}
	return nil
}
