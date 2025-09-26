// Package repository initDatabase.go
//
// 功能:
// - 连接到数据库，将db存储为包内变量
//
// 作者: LuckyQu
// 创建日期: 2025-09-24
// 修改日期: 2025-09-26

package repository

import (
	"backend/configs"
	"backend/internal/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// db 包内变量操作数据库
var db *gorm.DB

// InitDatabase 初始化连接到数据库，并将获取到的db存储到包内变量中
func InitDatabase() error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configs.Config.Mysql.Username,
		configs.Config.Mysql.Password,
		configs.Config.Mysql.Host,
		configs.Config.Mysql.Port,
		configs.Config.Mysql.DatabaseName,
	)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //关闭自动复数
		},
	})
	if err != nil {
		return err
	}
	db = database
	//初始化表结构
	//初始化用户表结构
	if !db.Migrator().HasTable(&model.User{}) {
		err = db.Migrator().CreateTable(&model.User{})
		if err != nil {
			return err
		}
	}
	//初始化题目表结构
	if !db.Migrator().HasTable(&model.Question{}) {
		err = db.Migrator().CreateTable(&model.Question{})
		if err != nil {
			return err
		}
	}
	return nil
}

// 对包内提供函数操作db
func getDB() *gorm.DB {
	return db
}
