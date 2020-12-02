package service

import (
	"Week02/dao"
	"database/sql"
	"errors"
	goerrors "github.com/pkg/errors"
	"log"
)

type Service struct {
	dao.Dao
}

func New()  *Service {
	return &Service{}
}

func (svr *Service) Update() (int,error) {
	_, err := svr.UpdateDb()
	if err != nil {
		if errors.Is(goerrors.Cause(err), sql.ErrNoRows) {
			log.Printf("[ERR]sql.ErrNoRows info：%T %v\n", goerrors.Cause(err), goerrors.Cause(err))
			log.Printf("stack：\n%+v\n", err)
			return 0,err
		}
	}
	return 120,nil;
}