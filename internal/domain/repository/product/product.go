package product

import postgre "github.com/ChanKachan/business-card-product/pkg/postgreWrapper"

type ProductRepository interface{}

type productRepository struct {
	postgres postgre.Postgre
}

func NewProductRepository(
	postgres postgre.Postgre,
) ProductRepository {
	return &productRepository{
		postgres: postgres,
	}
}
