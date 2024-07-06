package main

import (
	"log"

	"github.com/cjungo/cjuncms/misc"
	"github.com/cjungo/cjuncms/model"
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/rs/zerolog"
)

func main() {
	if err := cjungo.RunCommand[any](
		func(logger *zerolog.Logger, mysql *db.MySql, conf *db.MySqlConf) error {
			return mysql.AutoMigrate(
				&model.CjEmployeePermission{},
				&model.CjEmployee{},
				&model.CjMachineCPUTime{},
				&model.CjMachineVirtualMemory{},
				&model.CjOperation{},
				&model.CjPass{},
				&model.CjPermission{},
				&model.CjProjectEmployee{},
				&model.CjProject{},
				&model.CjScript{},
			)
		},
		db.LoadMySqlConfFormEnv,
		misc.ProvideMysqlForCmd(),
	); err != nil {
		log.Fatalln(err)
	}
}
