package gorm_demo

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log/slog"
)

type RebootLogs struct {
	gorm.Model
	Version      uint16 `gorm:"not null"`
	Status       uint   `gorm:"not null"`
	TimeInMillis uint64
	Flag         uint
}

func Sqlite_demo_3() {
	db, err := gorm.Open(sqlite.Open("./testdata/raptor.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate schema
	db.AutoMigrate(&RebootLogs{})

	//// Read
	//var product Product
	//db.First(&product, 1)                 // 根据整型主键查找
	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	////var log RebootLog
	//var logs []RebootLogs

	//// Time
	//db.Where("updated_at > ?", lastWeek).Find(&users)
	//// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';
	start := "2024-05-12 00:00:00"
	end := "2024-05-15 00:00:00"

	var count int64
	db.Model(&RebootLogs{}).Where("created_at > ? AND updated_at < ? AND status = ?", start, end, 0).Count(&count)
	slog.Info("打印查询结果", "Count", count)
	db.Model(&RebootLogs{}).Where("created_at > ? AND updated_at < ? AND status = ?", start, end, 1).Count(&count)
	slog.Info("打印查询结果", "Count", count)

	//db.Where("created_at > ? AND updated_at < ?", start, end).Find(&logs)
	//slog.Info("打印查询结果", "查询结果", logs)
	//for _, log := range logs {
	//	slog.Info("打印查询结果", "创建时间", log.CreatedAt)
	//	slog.Info("打印查询结果", "更新时间", log.UpdatedAt)
	//	slog.Info("打印查询结果", "创建时间", log.CreatedAt.UnixMilli())
	//	slog.Info("打印查询结果", "更新时间", log.UpdatedAt.UnixMilli())
	//
	//	db.Delete(&log)
	//}

	//db.Select("created_at", "updated_at").Where("created_at > ? AND updated_at < ?", start, end).Find(&logs)
	//
	//// Print the selected fields
	//fmt.Printf("Selected fields: %+v\n", logs)

	//// BETWEEN
	//db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
	//// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

}
