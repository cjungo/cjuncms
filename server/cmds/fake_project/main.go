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
		func(logger *zerolog.Logger, mysql *db.MySql) {
			count := 1000
			ms := make([]model.CjProject, count)
			for i := range count {
				ms[i] = model.CjProject{
					Name: faker.ChineseName(),
				}
			}

		},
		db.LoadMySqlConfFormEnv,
		misc.ProvideMysql(),
	); err != nil {
		log.Fatalln(err)
	}
}
