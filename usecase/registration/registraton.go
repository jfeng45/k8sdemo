package registration

import (
	"github.com/jfeng45/k8sdemo/dataservice"
	"github.com/jfeng45/k8sdemo/model"
	"github.com/pkg/errors"
)

type UseCase struct {
	UserDataInterface dataservice.UserDataInterface
}

func (uuc *UseCase)Register(user *model.User) (bool, error) {
	isD, err := uuc.isDuplicate(user.Name)
	if err != nil {
		return false, err
	}
	if isD {
		return false, errors.New("duplicate user for "+ user.Name)
	}
	 uuc.UserDataInterface.Insert(user)
	//if rowCount != 1 {
	//	return false, errors.New("Register failed and the number f rows affected are  "+ string(rowCount))
	//}
	return true, nil
}

func (uuc *UseCase) ListUser () ([]model.User, error){
	return uuc.UserDataInterface.FindAll()
	//return []model.User{}, nil
}
func (uuc *UseCase) isDuplicate(name string) (bool, error) {
	return uuc.UserDataInterface.FindByName(name)
	//return false
}

