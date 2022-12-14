package database

import (
	"setup/environment"
	"setup/helpers"
	models "setup/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*
If you are using Windows
	- Download the GCC from here http://tdm-gcc.tdragon.net/download
	- Install GCC
*/

func (d *Sqlite) main() *gorm.DB {
	d.db, err = gorm.Open(sqlite.Open("./database/"+helpers.GoDotEnvVariable(environment.DB_DATABASE)+".db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	d.db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err.Error())
	}
	return d.db
}
