package dao

import (
	"database/sql"
	"github.com/pkg/errors"

)
type Dao struct {

}

func New() *Dao {
	return &Dao{}
}

func (dao *Dao) UpdateDb() (int, error) {
	return 0, errors.Wrap(sql.ErrNoRows, "dao update DB error")
}

