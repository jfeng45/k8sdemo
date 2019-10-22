package usecase

import (
	"github.com/jfeng45/k8sdemo/model"
)

type RegistrationUseCaseInterface interface {
	Register(user *model.User) (bool, error)
	ListUser () ([]model.User, error)
}


