package user

import (
	"auth/model"
	"auth/tool/hash"
	"auth/tool/parser"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type defaultDAO struct {
	db *gorm.DB
}

func NewDefaultDAO(db *gorm.DB) *defaultDAO {
	return &defaultDAO{
		db: db,
	}
}

func (d *defaultDAO) Insert(u *model.Auth) (result *model.Auth, err error) {
	if u.UserPw, err = hash.BcryptGenerate(u.UserPw, bcrypt.DefaultCost); err != nil {
		err = BcryptGenerateError
		return
	}
	u.Status = CreatePending

	r := d.db.Create(u)
	if r.Error == nil { result = r.Value.(*model.Auth); return }

	code, err := parser.DBErrorParse(r.Error.Error())
	if err != nil { err = parser.InvalidError; return }

	switch code {
	case IdDuplicateErrorCode:
		err = IdDuplicateError
	case DataTooLongErrorCode:
		err = DataLengthOverError
	default:
		err = r.Error
	}
	return
}

func (d *defaultDAO) Commit() *gorm.DB {
	return d.db.Commit()
}

func (d *defaultDAO) Rollback() *gorm.DB {
	return d.db.Rollback()
}