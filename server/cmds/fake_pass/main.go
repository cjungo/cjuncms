package main

import (
	"log"

	"github.com/cjungo/cjuncms/misc"
	"github.com/cjungo/cjuncms/model"
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/go-faker/faker/v4"
	"github.com/rs/zerolog"
)

func main() {
	if err := cjungo.RunCommand[any](
		func(logger *zerolog.Logger, mysql *db.MySql) error {
			count := 1000
			rows := make([]model.CjPass, count)
			for i := range count {
				title := faker.ChineseName()
				rows[i] = model.CjPass{
					Type:    1,
					Title:   title,
					Host:    "127.0.0.1",
					Port:    22,
					Content: "123456",
				}
			}
			if err := mysql.CreateInBatches(&rows, 100).Error; err != nil {
				return err
			}
			return nil
		},
		db.LoadMySqlConfFormEnv,
		misc.ProvideMysqlForWeb(),
	); err != nil {
		log.Fatalln(err)
	}
}
