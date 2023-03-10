package database

import (
	"fmt"
	"waysbeans/models"
	"waysbeans/pkg/mysql"
)

func RunMigration() {
	err := mysql.ConnDB.AutoMigrate(&models.Product{})

	if err != nil {
		fmt.Println(err)
		panic("migration error")
	}

	fmt.Println("migration success")
}
