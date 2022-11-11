package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql() *gorm.DB {
	mysqlConfig := mysql.Config{
		DSN:                       Dsn(), // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(100)
		sqlDB.SetMaxOpenConns(100)
		return db
	}
}

func Dsn() string {
	return "root" + ":" + "DCone2020!" + "@tcp(" + "localhost" + ":" + "3306" + ")/" + "casbin"
}
