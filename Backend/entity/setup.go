package entity

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
var db *gorm.DB

func DB() *gorm.DB {
	return db
}

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n=============================\n", sql)
}


func SetupDatabase() {
  database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{
    Logger: SqlLogger{},
  })
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		//01 User
		&User{},
		&UserType{},
		// 05 ระบบโภชนาการ
		&Manage{},
		&Nutrition{},

	)

	//05 Manage ข้อมูลที่ต้องเตรียมไว้ก่อน
	database.Create(&Nutrition{Type: "กำหนดเอง",Receive: 0, Detail: "มีการจัดโภชนาการตามแพทย์เห็นสมควร"})
	database.Create(&Nutrition{Type: "อาหารอ่อน, นิ่ม",Receive: 2000, Detail: "ข้าวต้ม, นม, มะม่าง"})
	database.Create(&Nutrition{Type: "อาหารที่มีการเคี๊ยวหน่อย",Receive: 2200, Detail: "ไข่ต้ม, แตงกวา, ข้าวผัด, นม, มะม่วง"})
	
  	db = database
}