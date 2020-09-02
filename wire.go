package main

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"go-gin-gorm-di/product"
)

func initProductAPI(db *gorm.DB) product.ProductAPI {
	wire.Build(product.ProvideProductRepostiory, product.ProvideProductService, product.ProvideProductAPI)

	return product.ProductAPI{}
}
