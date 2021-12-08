package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const MYDB_DSN = "root:123456@tcp(192.168.232.100:3306)/sql_test?charset=utf8mb4&parseTime=True&loc=Local"

func InitDB() (*gorm.DB, error) {
	sqlDb, err := gorm.Open(mysql.Open(MYDB_DSN), &gorm.Config{})
	if err != nil {
		fmt.Printf("init mysql db failed,err:%v", err)
		return nil, err
	}
	return sqlDb, nil
}
