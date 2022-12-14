// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/acxe21/crud-app.git/handlers"
	"github.com/acxe21/crud-app.git/repository"
	"github.com/acxe21/crud-app.git/services"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func ProductHandler(db *gorm.DB) handler.ProductHandler {
	testProductRepository := repository.ProvideProductRepository(db)
	testProductService := service.ProvideProductService(testProductRepository)
	productHandler := handler.ProvideProductHandle(testProductService)
	return productHandler
}
