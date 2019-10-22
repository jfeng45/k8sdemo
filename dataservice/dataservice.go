package dataservice

import (
	"github.com/jfeng45/k8sdemo/model"
)

type UserDataInterface interface {
	Remove(err error, id int64, affect int64)
	Find(id int) error
	FindByName(name string) (bool, error)
	FindAll() ([]model.User, error)
	Update(user *model.User) (error, int64)
	Insert(user *model.User) int64
	Close() error
}





