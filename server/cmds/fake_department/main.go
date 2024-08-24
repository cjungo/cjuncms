package main

import (
	"fmt"
	"log"
	"time"

	"github.com/cjungo/cjuncms/misc"
	"github.com/cjungo/cjuncms/model"
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/rs/zerolog"
)

func main() {
	if err := cjungo.RunCommand[any](
		func(logger *zerolog.Logger, mysql *db.MySql) error {
			lv1Count := 40
			lv2Count := 40
			lv3Count := 40
			now := time.Now()
			rows := make([]model.CjDepartment, lv1Count+lv1Count*lv2Count+lv1Count*lv2Count*lv3Count)
			for i := range lv1Count {
				lv1ID := uint32(1000000 + 100000*(i+1))
				rows[i] = model.CjDepartment{
					ID:       lv1ID,
					Level:    1,
					ParentID: 0,
					Title:    fmt.Sprintf("[V1]部门 %d", i+1),
					CreateAt: now,
				}
				for j := range lv2Count {
					lv2ID := lv1ID + uint32((j+1)*1000)
					rows[lv1Count+i*lv2Count+j] = model.CjDepartment{
						ID:       lv2ID,
						ParentID: lv1ID,
						Level:    2,
						Title:    fmt.Sprintf("[V2]部门 %d", lv2ID),
						CreateAt: now,
					}

					for k := range lv3Count {
						lv3ID := lv2ID + uint32(k+1)
						rows[lv1Count+lv1Count*lv2Count+i*lv2Count*lv3Count+j*lv3Count+k] = model.CjDepartment{
							ID:       lv3ID,
							ParentID: lv2ID,
							Level:    3,
							Title:    fmt.Sprintf("[V3]部门 %d", lv3ID),
							CreateAt: now,
						}
					}
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
