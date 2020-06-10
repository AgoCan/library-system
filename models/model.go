package model
// https://gorm.io/zh_CN/
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"library-system/config"
)
// DB db handler
var DB *gorm.DB
// InitMysql 初始化数据库
func InitMysql() (err error) {
	DB, err = gorm.Open("mysql", config.MysqlConnect)
	if err != nil {
		return err
	}
	// 创建数据库
	migrate()
	// 尝试连接
	err = DB.DB().Ping()
	return err
}
// Close 关闭数据库
func Close()(){
	err := DB.Close()
	if err != nil {
		fmt.Println(err)
	}
	return
}
func migrate(){
	userInfoMigrate()
	blogMigrate()
}