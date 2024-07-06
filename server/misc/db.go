package misc

import (
	"github.com/cjungo/cjuncms/entity"
	"github.com/cjungo/cjungo/db"
	"gorm.io/gorm"
)

func ProvideMysqlWith(action func(mysql *db.MySql) error) db.MySqlProvide {
	return db.NewMySqlHandle(func(mysql *db.MySql) error {
		entity.Use(mysql.DB)
		return action(mysql)
	})
}

func ProvideMysqlForCmd() db.MySqlProvide {
	return ProvideMysqlWith(func(mysql *db.MySql) error {
		return nil
	})
}

func ProvideMysqlForWeb() db.MySqlProvide {
	return ProvideMysqlWith(func(mysql *db.MySql) error {
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
