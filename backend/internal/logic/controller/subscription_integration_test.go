//go:build integration

package controller

import (
	"DoramaSet/internal/container"
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/repository/postgres"
	"context"
	"errors"
	"testing"
)

func TestSubscriptionController_SubscribeUserIntegration(t *testing.T) {
	dbContainer, db, err := container.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer dbContainer.Terminate(context.Background())
	repo := postgres.NewSubscriptionRepo(db)
	pr := postgres.NewPictureRepo(db)
	er := postgres.NewEpisodeRepo(db)
	dr := postgres.NewDoramaRepo(db, pr, er)
	lr := postgres.NewListRepo(db, dr)
	urepo := postgres.NewUserRepo(db, repo, lr)
	pointC := PointsController{urepo}
	userC := UserController{repo: urepo, pc: &pointC, secretKey: "qwerty"}
	token := getToken(model.User{Username: "test"}, "qwerty")
	type fields struct {
		repo  repository.ISubscriptionRepo
		urepo repository.IUserRepo
		pc    controller.IPointsController
		uc    controller.IUserController
	}
	type args struct {
		token string
		id    int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		check   func(user repository.IUserRepo) error
	}{
		{
			name: "subscribe user",
			fields: fields{
				repo:  repo,
				urepo: urepo,
				pc:    &pointC,
				uc:    &userC,
			},
			args:    args{token: token, id: 2},
			wantErr: false,
			check: func(userRepo repository.IUserRepo) error {
				user, err := userRepo.GetUser("test")
				if err != nil {
					return err
				}
				if user.Sub.Id != 2 {
					return errors.New("error")
				}
				return nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SubscriptionController{
				repo:  tt.fields.repo,
				urepo: tt.fields.urepo,
				pc:    tt.fields.pc,
				uc:    tt.fields.uc,
			}
			if err := s.SubscribeUser(tt.args.token, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("SubscribeUserIntegration() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := tt.check(urepo); (err != nil) != tt.wantErr {
				t.Errorf("SubscribeUserIntegration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
