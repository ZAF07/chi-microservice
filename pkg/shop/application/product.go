// THIS PACKAGE IS LIKE A CONTROLLER MODULE
package application

import (
	"github.com/ZAF07/chi-microservice/pkg/common/price"
)

type productReadModel interface {
	AllProducts() ([]products.Product, err)
}

type ProductsService struct {
	repo products.Repository
	repoModel productReadModel
}

func NewProductsService() ProductsService {

}

func (s ProductsService) AllProducts() () {

}

type AddProductCommand struct {
	ID 						string
	Name 					string
	Description 	string
	PriceCents 		uint
	PriceCurrency string
}

func (s ProductsService) AddProduct(cmd AllProductCommand) error {
	price.NewPrice(cmd.PriceCents, cmd.PriceCurrency)

	products.NewProduct(products.ID(cmd.ID), cmd.Name, cmd.Description, price)

	s.repo.Save
}