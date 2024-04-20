package router

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	m "hexTest/model"
)

func NewDB() (*gorm.DB, error) {
	InitConfig()

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.name"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if viper.GetBool("db.auto_migration") {
		db.AutoMigrate(&m.Beer{}, &m.User{}, &m.Company{}, &m.Distributer{}, &m.DistributerBeer{})
	}

	return db, nil
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
