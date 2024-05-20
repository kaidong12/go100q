package gorm_demo

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log/slog"
	"time"
)

type RebootLog struct {
	gorm.Model
	Status       uint
	Version      uint16
	TimeInMillis uint32
	Flag         uint
}

func recordLog(db *gorm.DB, log *RebootLog) {
	var last RebootLog
	result := db.First(&last, RebootLog{Flag: 1})
	if result.RowsAffected == 0 {
		slog.Info("Data table is empty", "Create a new record: ", log)
		db.Create(&log)
		return
	} else {
		slog.Info("===================", "Current record is: ", last)
		if last.Status == log.Status {
			slog.Info("Update Timeins", "New Timeins: ", log.TimeInMillis)
			db.Model(&last).Update("Timeins", log.TimeInMillis)
		} else {
			//db.Model(&last).Updates(RebootLog{Timeins: last.TimeInMillis + 60000, Flag: 0})
			db.Model(&last).Updates(map[string]interface{}{"Timeins": last.TimeInMillis + 60000, "Flag": 0})

			slog.Info("Print new record", "New record: ", log)
			log.Flag = 1
			log.Version = last.Version + 1
			slog.Info("Insert new record", "New record: ", log)
			db.Create(&log)
		}
	}
}

func Sqlite_demo_2() {
	db, err := gorm.Open(sqlite.Open("./testdata/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate schema
	db.AutoMigrate(&RebootLog{})

	// Create
	//recordLog(db, &RebootLog{Status: 1, Version: 1, Flag: 1, Timeins: 0})
	//time.Sleep(time.Second * 1)
	//recordLog(db, &RebootLog{Status: 1, Timeins: 10})
	//time.Sleep(time.Second * 1)
	//
	//recordLog(db, &RebootLog{Status: 0, Timeins: 0})
	//time.Sleep(time.Second * 1)
	//recordLog(db, &RebootLog{Status: 0, Timeins: 10})
	//time.Sleep(time.Second * 1)

	recordLog(db, &RebootLog{Status: 0, TimeInMillis: 20000})
	time.Sleep(time.Second * 1)
	recordLog(db, &RebootLog{Status: 0, TimeInMillis: 40000})
	time.Sleep(time.Second * 1)

	recordLog(db, &RebootLog{Status: 1, TimeInMillis: 0})
	time.Sleep(time.Second * 1)
	recordLog(db, &RebootLog{Status: 1, TimeInMillis: 10000})
	time.Sleep(time.Second * 1)

	//// Read
	//var product Product
	//db.First(&product, 1)                 // 根据整型主键查找
	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	//var log RebootLog
	var logs []RebootLog

	//// Time
	//db.Where("updated_at > ?", lastWeek).Find(&users)
	//// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';
	start := "2024-05-04 00:00:00"
	end := "2024-05-05 00:00:00"
	db.Where("created_at > ? AND updated_at < ?", start, end).Find(&logs)
	slog.Info("打印查询结果", "查询结果", logs)
	for _, log := range logs {
		slog.Info("打印查询结果", "创建时间", log.CreatedAt)
		slog.Info("打印查询结果", "更新时间", log.UpdatedAt)
		slog.Info("打印查询结果", "创建时间", log.CreatedAt.UnixMilli())
		slog.Info("打印查询结果", "更新时间", log.UpdatedAt.UnixMilli())

		db.Delete(&log)
	}

	//db.Select("created_at", "updated_at").Where("created_at > ? AND updated_at < ?", start, end).Find(&logs)
	//
	//// Print the selected fields
	//fmt.Printf("Selected fields: %+v\n", logs)

	//// BETWEEN
	//db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
	//// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

	//// Update - 将 product 的 price 更新为 200
	//db.Model(&product).Update("Price", 200)
	//// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - 删除 product
	//db.Delete(&product, 1)

}
