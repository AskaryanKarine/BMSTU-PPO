package postgres

import (
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/errors"
	"DoramaSet/internal/logic/model"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type DoramaRepo struct {
	db      *gorm.DB
	picRepo repository.IPictureRepo
	epRepo  repository.IEpisodeRepo
}

type doramaModel struct {
	ID          int
	Name        string
	Description string
	ReleaseYear int
	Status      string
	Genre       string
}

func NewDoramaRepo(db *gorm.DB, PR repository.IPictureRepo, ER repository.IEpisodeRepo) *DoramaRepo {
	return &DoramaRepo{db, PR, ER}
}

func (d *DoramaRepo) GetList() ([]model.Dorama, error) {
	var (
		resDB []doramaModel
		res   []model.Dorama
	)
	result := d.db.Table("dorama_set.dorama").Find(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	if len(resDB) == 0 {
		return nil, nil
	}

	for _, r := range resDB {
		ep, err := d.epRepo.GetList(r.ID)
		if err != nil {
			return nil, fmt.Errorf("getListEp: %w", err)
		}
		photo, err := d.picRepo.GetListDorama(r.ID)
		if err != nil {
			return nil, fmt.Errorf("getListDorama: %w", err)
		}
		tmp := model.Dorama{
			Id:          r.ID,
			Name:        r.Name,
			Description: r.Description,
			Genre:       r.Genre,
			ReleaseYear: r.ReleaseYear,
			Episodes:    ep,
			Posters:     photo,
		}
		res = append(res, tmp)
	}
	return res, nil
}

func (d *DoramaRepo) GetListName(name string) ([]model.Dorama, error) {
	var (
		resDB []doramaModel
		res   []model.Dorama
	)
	likeStr := "%" + strings.TrimRight(name, "\r\n") + "%"
	result := d.db.Table("dorama_set.dorama").Where("name like ?", likeStr).Find(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	if len(resDB) == 0 {
		return nil, fmt.Errorf("db: %w", errors.ErrorDontExistsInDB)
	}

	for _, r := range resDB {
		ep, err := d.epRepo.GetList(r.ID)
		if err != nil {
			return nil, fmt.Errorf("getListEp: %w", err)
		}
		photo, err := d.picRepo.GetListDorama(r.ID)
		if err != nil {
			return nil, fmt.Errorf("getListDorama: %w", err)
		}
		tmp := model.Dorama{
			Id:          r.ID,
			Name:        r.Name,
			Description: r.Description,
			Genre:       r.Genre,
			ReleaseYear: r.ReleaseYear,
			Episodes:    ep,
			Posters:     photo,
		}
		res = append(res, tmp)
	}
	return res, nil
}

func (d *DoramaRepo) GetDorama(id int) (*model.Dorama, error) {
	var (
		resDB doramaModel
		res   model.Dorama
	)
	result := d.db.Table("dorama_set.dorama").Where("id = ?", id).Take(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	ep, err := d.epRepo.GetList(resDB.ID)
	if err != nil {
		return nil, fmt.Errorf("getListEp: %w", err)
	}
	photo, err := d.picRepo.GetListDorama(resDB.ID)
	if err != nil {
		return nil, fmt.Errorf("getListDorama: %w", err)
	}
	res = model.Dorama{
		Id:          resDB.ID,
		Name:        resDB.Name,
		Description: resDB.Description,
		Status:      resDB.Status,
		Genre:       resDB.Genre,
		ReleaseYear: resDB.ReleaseYear,
		Episodes:    ep,
		Posters:     photo,
	}

	return &res, nil
}

func (d *DoramaRepo) CreateDorama(dorama model.Dorama) (int, error) {
	m := doramaModel{
		Name:        dorama.Name,
		Description: dorama.Description,
		ReleaseYear: dorama.ReleaseYear,
		Status:      dorama.Status,
		Genre:       dorama.Genre,
	}
	result := d.db.Table("dorama_set.dorama").Create(&m)
	if result.Error != nil {
		return -1, fmt.Errorf("db: %w", result.Error)
	}
	return m.ID, nil
}

func (d *DoramaRepo) UpdateDorama(dorama model.Dorama) error {
	m := doramaModel{
		ID:          dorama.Id,
		Name:        dorama.Name,
		Description: dorama.Description,
		ReleaseYear: dorama.ReleaseYear,
		Status:      dorama.Status,
		Genre:       dorama.Genre,
	}
	result := d.db.Table("dorama_set.dorama").Save(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (d *DoramaRepo) DeleteDorama(id int) error {
	result := d.db.Table("dorama_set.dorama").Where("id = ?", id).Delete(&doramaModel{})
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (d *DoramaRepo) AddStaff(idD, idS int) error {
	m := struct {
		IdDorama, IdStaff int
	}{idD, idS}
	result := d.db.Table("dorama_set.doramastaff").Create(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (d *DoramaRepo) GetListByListId(idL int) ([]model.Dorama, error) {
	var (
		resDB []struct {
			IdDorama int
			IdList   int
		}
		res []model.Dorama
	)
	result := d.db.Table("dorama_set.listdorama").Where("id_list = ?", idL).Find(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	for _, r := range resDB {
		dorama, err := d.GetDorama(r.IdDorama)
		if err != nil {
			return nil, fmt.Errorf("getDorama: %w", err)
		}
		res = append(res, *dorama)
	}
	return res, nil
}
