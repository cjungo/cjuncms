package misc

import (
	"github.com/cjungo/cjuncms/entity"
	"github.com/cjungo/cjungo/db"
	"gorm.io/gorm"
)

func ProvideMysql() db.MySqlProvide {
	return db.NewMySqlHandle(func(mysql *db.MySql) error {
		entity.Use(mysql.DB)
		return mysql.Transaction(func(tx *gorm.DB) error {
			if err := EnsurePermissions(tx); err != nil {
				return err
			}

			if err := EnsureAdmin(tx); err != nil {
				return err
			}

			return nil
		})
	})
}
