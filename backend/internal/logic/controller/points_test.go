package controller

import (
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/repository/mocks"
	"errors"
	"github.com/sirupsen/logrus"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
)

func TestEarnPointForLogin(t *testing.T) {
	mc := minimock.NewController(t)
	userFev := model.User{RegData: time.Date(2000, time.February, 29, 0, 0, 0, 0, time.UTC)}
	userReg := model.User{RegData: time.Now()}
	userReg2 := model.User{RegData: time.Now().AddDate(0, 1, 0)}
	testsTable := []struct {
		name  string
		fl    PointsController
		arg   model.User
		isNeg bool
	}{
		{
			name: "successful result",
			fl: PointsController{
				repo: mocks.NewIUserRepoMock(mc).GetUserMock.Return(&model.User{}, nil).UpdateUserMock.Return(nil),
			},
			arg:   model.User{},
			isNeg: false,
		},
		{
			name: "update error",
			fl: PointsController{
				repo: mocks.NewIUserRepoMock(mc).GetUserMock.Return(&model.User{}, nil).UpdateUserMock.Return(errors.New("error")),
			},
			arg:   model.User{},
			isNeg: true,
		},
		{
			name: "29 february registration",
			fl: PointsController{
				repo: mocks.NewIUserRepoMock(mc).GetUserMock.Return(&userFev, nil).UpdateUserMock.Return(nil),
			},
			arg:   userFev,
			isNeg: false,
		},
		{
			name: "at the moment registration",
			fl: PointsController{
				repo: mocks.NewIUserRepoMock(mc).GetUserMock.Return(&userReg, nil).UpdateUserMock.Return(nil),
			},
			arg:   userReg,
			isNeg: false,
		},
		{
			name: "successful registration",
			fl: PointsController{
				repo: mocks.NewIUserRepoMock(mc).GetUserMock.Return(&userReg2, nil).UpdateUserMock.Return(nil),
			},
			arg:   userReg2,
			isNeg: false,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := PointsController{
				repo: testCase.fl.repo,
				log:  &logrus.Logger{},
			}
			err := dc.EarnPointForLogin(&testCase.arg)
			if (err != nil) != testCase.isNeg {
				t.Errorf("EarnPointForLogin() error = %v, expect = %v", err, testCase.isNeg)
			}
		})
	}
}

func TestPurgePoint(t *testing.T) {
	mc := minimock.NewController(t)
	type argument struct {
		username model.User
		point    int
	}
	testsTable := []struct {
		name  string
		fl    PointsController
		arg   argument
		isNeg bool
	}{
		{
			name: "successful result",
			fl: PointsController{
				repo: mocks.NewIUserRepoMock(mc).GetUserMock.Return(&model.User{Points: 10}, nil).UpdateUserMock.Return(nil),
			},
			arg:   argument{model.User{Points: 10}, 1},
			isNeg: false,
		},
		{
			name: "update error",
			fl: PointsController{
				repo: mocks.NewIUserRepoMock(mc).GetUserMock.Return(&model.User{Points: 10}, nil).UpdateUserMock.Return(errors.New("error")),
			},
			arg:   argument{model.User{Points: 10}, 1},
			isNeg: true,
		},
		{
			name: "negative balance",
			fl: PointsController{
				repo: mocks.NewIUserRepoMock(mc).GetUserMock.Return(&model.User{}, nil).UpdateUserMock.Return(errors.New("error")),
			},
			arg:   argument{model.User{}, 1},
			isNeg: true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := PointsController{
				repo: testCase.fl.repo,
				log:  &logrus.Logger{},
			}
			err := dc.PurgePoint(&testCase.arg.username, testCase.arg.point)
			if (err != nil) != testCase.isNeg {
				t.Errorf("PurgePoint() error = %v, expect = %v", err, testCase.isNeg)
			}
		})
	}
}

func TestEarnPoint(t *testing.T) {
	mc := minimock.NewController(t)
	type argument struct {
		username model.User
		point    int
	}
	testsTable := []struct {
		name  string
		fl    PointsController
		arg   argument
		isNeg bool
	}{
		{
			name: "successful result",
			fl: PointsController{
				repo: mocks.NewIUserRepoMock(mc).GetUserMock.Return(&model.User{Points: 10}, nil).UpdateUserMock.Return(nil),
			},
			arg:   argument{model.User{Points: 10}, 1},
			isNeg: false,
		},
		{
			name: "update error",
			fl: PointsController{
				repo: mocks.NewIUserRepoMock(mc).GetUserMock.Return(&model.User{Points: 10}, nil).UpdateUserMock.Return(errors.New("error")),
			},
			arg:   argument{model.User{Points: 10}, 1},
			isNeg: true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := PointsController{
				repo: testCase.fl.repo,
				log:  &logrus.Logger{},
			}
			err := dc.EarnPoint(&testCase.arg.username, testCase.arg.point)
			if (err != nil) != testCase.isNeg {
				t.Errorf("EarnPoint() error = %v, expect = %v", err, testCase.isNeg)
			}
		})
	}
}
