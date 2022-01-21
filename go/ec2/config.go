package main

import (
	"fmt"
	"github.com/copydataai/backend-practice/go/ec2/models/"
	config "github.com/spf13/viper"
	"gorm.io/gorm"
	"os"
)

func migrate(database *gorm.DB) {
	database.AutoMigrate(&models.Reviews{})
	database.AutoMigrate(&models.Products{})
	database.AutoMigrate(&models.ProductsType{})
	database.AutoMigrate(&models.Users{})
}
