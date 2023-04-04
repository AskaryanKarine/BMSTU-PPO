package controller

import (
	"DoramaSet/internal/interfaces"
	"DoramaSet/internal/logic/model"
	"errors"
	"fmt"
)

type StaffController struct {
	repo interfaces.IStaffRepo
	uc   interfaces.IUserController
}

func (s *StaffController) GetList() ([]model.Staff, error) {
	res, err := s.repo.GetList()
	if err != nil {
		return nil, fmt.Errorf("getList: %w", err)
	}
	return res, nil
}

func (s *StaffController) GetListByName(name string) ([]model.Staff, error) {
	return s.repo.GetListName(name)
}

func (s *StaffController) GetListByDorama(idD int) ([]model.Staff, error) {
	res, err := s.repo.GetListDorama(idD)
	if err != nil {
		return nil, fmt.Errorf("getListByDorama: %w", err)
	}
	return res, nil
}

func (s *StaffController) CreateStaff(token string, record model.Staff) error {
	user, err := s.uc.AuthByToken(token)
	if err != nil {
		return fmt.Errorf("createStaff: %w", err)
	}

	if !user.IsAdmin {
		return errors.New("createStaff: low level of access")
	}

	err = s.repo.CreateStaff(record)
	if err != nil {
		return fmt.Errorf("createStaff: %w", err)
	}
	return nil
}

func (s *StaffController) UpdateStaff(token string, record model.Staff) error {
	user, err := s.uc.AuthByToken(token)
	if err != nil {
		return fmt.Errorf("updateStaff: %w", err)
	}

	if !user.IsAdmin {
		return errors.New("updateStaff: low level of access")
	}
	err = s.repo.UpdateStaff(record)
	if err != nil {
		return fmt.Errorf("updateStaff: %w", err)
	}
	return nil
}
