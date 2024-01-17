package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	//MustConnect--> 如果没有连接上就panic掉
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_con"))
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_con"))

	//建表操作
	err = CreateUserTale(db)
	if err != nil {
		zap.L().Error("CreateUserTale(db)failed,err:%v", zap.Error(err))
		return err
	}

	return
}

func CreateUserTale(db *sqlx.DB) (err error) {
	sqlStr := `	CREATE TABLE IF NOT EXISTS users (
			id BIGINT AUTO_INCREMENT PRIMARY KEY,
			user_id BIGINT NOT NULL ,
			username VARCHAR(24) NOT NULL,
			password VARCHAR(96) NOT NULL,
			email VARCHAR(32) NOT NULL,
			create_time timestamp default CURRENT_TIMESTAMP NULL,
			update_time timestamp default NULL ON UPDATE CURRENT_TIMESTAMP)`
	_, err = db.Exec(sqlStr)
	return err
}