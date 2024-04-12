package tidbClient

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Get the env variables from GitHub workflow
func getDSN() string {
	tidbUser := "root"
	tidbPassword := ""
	tidbHost := "127.0.0.1"
	tidbPort := "4000"

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/test?charset=utf8mb4",
		tidbUser, tidbPassword, tidbHost, tidbPort)
}

func createDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(getDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	return db
}

func GetTiDBVersion() string {
	db := createDB()
	var version string
	db.Raw("SELECT VERSION()").Scan(&version)

	return version
}
