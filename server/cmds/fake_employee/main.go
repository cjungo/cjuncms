package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/cjungo/cjuncms/misc"
	"github.com/cjungo/cjuncms/model"
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/cjungo/cjungo/ext"
	"github.com/go-faker/faker/v4"
	"github.com/rs/zerolog"
)

func main() {
	if err := cjungo.RunCommand[any](
		func(logger *zerolog.Logger, mysql *db.MySql) error {
			count := 1000
			baseDay, err := time.Parse("2006-02-01", "2000-01-01")
			if err != nil {
				return nil
			}
			rows := make([]model.CjEmployee, count)
			for i := range count {
				fullname := faker.ChineseName()
				nickname := faker.ChineseName()
				birthday := baseDay.Add(time.Hour * 24 * time.Duration(rand.Int63n(9999)))
				rows[i] = model.CjEmployee{
					Username:  faker.Username(),
					Password:  ext.Sha256("123456"),
					Jobnumber: fmt.Sprintf("1%05d", i),
					Fullname:  &fullname,
					Nickname:  &nickname,
					Birthday:  &birthday,
				}
			}
			if err := mysql.CreateInBatches(&rows, 100).Error; err != nil {
				return err
			}
			return nil
		},
		db.LoadMySqlConfFormEnv,
		misc.ProvideMysql(),
	); err != nil {
		log.Fatalln(err)
	}
}
