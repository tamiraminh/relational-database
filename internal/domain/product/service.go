package product

import (
	"github.com/evermos/boilerplate-go/configs"
	"github.com/evermos/boilerplate-go/event/producer"
	"github.com/evermos/boilerplate-go/shared/failure"
	"github.com/gofrs/uuid"
)


type ProductService interface {
	Create(requestFormat ProductRequestFormat) (product Product, err error)
	Update(id uuid.UUID, requestFormat ProductRequestFormat) (product Product, err error)
	SoftDelete(id uuid.UUID, userID uuid.UUID) (product Product, err error)
	HardDelete(id uuid.UUID, userID uuid.UUID) (product Product, err error)
	ReadPagination(limit int, offset int) (products []Product, err error)
	ReadStatusSorted() (products []ProductStatus, err error) 

}


type ProductServiceImpl struct {
	ProductRepository ProductRepository
	Config        *configs.Config
}


func ProvideProductServiceImpl(productRepository ProductRepository, producer producer.Producer, config *configs.Config) *ProductServiceImpl {
	s := new(ProductServiceImpl)
	s.ProductRepository = productRepository
	s.Config = config

	return s
}


func (s *ProductServiceImpl) Create(requestFormat ProductRequestFormat) (product Product, err error) {
	product, err = product.NewFromRequestFormat(requestFormat)
	if err != nil {
		return
	}

	if err != nil {
		return product, failure.BadRequest(err)
	}

	err = s.ProductRepository.Create(product)

	if err != nil {
		return
	}


	return
}


func (s *ProductServiceImpl) Update(id uuid.UUID, requestFormat ProductRequestFormat) (product Product, err error) {
	product, err = s.ProductRepository.ResolveByID(id)
	if err != nil {
		return
	}

	err = product.Update(requestFormat, requestFormat.UserId)
	if err != nil {
		return
	}

	err = s.ProductRepository.Update(product)
	return
}


func (s *ProductServiceImpl) SoftDelete(id uuid.UUID, userID uuid.UUID) (product Product, err error) {
	product, err = s.ProductRepository.ResolveByID(id)
	if err != nil {
		return
	}

	variants, err := s.ProductRepository.ResolveVariantsByProductIDs([]uuid.UUID{product.Id})
	if err != nil {
		return product, err
	}

	product.AttachVariants(variants)

	err = product.SoftDelete(userID)
	if err != nil {
		return
	}

	err = s.ProductRepository.Update(product)
	

	return
}


func (s *ProductServiceImpl) HardDelete(id uuid.UUID, userID uuid.UUID) (product Product, err error) {
	product, err = s.ProductRepository.ResolveByID(id)
	if err != nil {
		return
	}

	variants, err := s.ProductRepository.ResolveVariantsByProductIDs([]uuid.UUID{product.Id})
	if err != nil {
		return product, err
	}

	product.AttachVariants(variants)


	err = s.ProductRepository.HardDelete(id)
	if err != nil {
		return product, failure.BadRequest(err)
	}

	return
}


func (s *ProductServiceImpl) ReadPagination(limit int, page int) (products []Product,err error){
	products, err = s.ProductRepository.ReadPagination(limit, page)
	if err != nil {
		return
	}

	return
}

func (s *ProductServiceImpl) ReadStatusSorted() (products []ProductStatus,err error){
	products, err = s.ProductRepository.ReadStatusSorted()
	if err != nil {
		return
	}

	return
}