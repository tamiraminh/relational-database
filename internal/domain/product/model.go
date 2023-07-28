package product

import (
	"encoding/json"
	"time"

	"github.com/evermos/boilerplate-go/shared/failure"
	"github.com/evermos/boilerplate-go/shared/nuuid"
	"github.com/gofrs/uuid"
	"github.com/guregu/null"
)

type Product struct {
	Id   			uuid.UUID		`db:"id"`
	BrandId 		uuid.UUID		`db:"brandId"`
	Name			string	 		`db:"name"`
	CreatedAt       time.Time  		`db:"createdAt" `
	CreatedBy     	uuid.UUID   	`db:"createdBy"`
	UpdatedAt       null.Time   	`db:"updatedAt"`
	UpdatedBy     	nuuid.NUUID		`db:"updatedBy"`
	DeletedAt       null.Time   	`db:"deletedAt"`
	DeletedBy     	nuuid.NUUID		`db:"deletedBy"`
	Variants        []Variant   	`db:"-" validate:"required,dive,required"`
}

func (p *Product) AttachVariants(variants []Variant) Product {
	for _, item := range variants {
		if item.ProductId == p.Id {
			p.Variants = append(p.Variants, item)
		}
	}
	return *p
}


func (p *Product) IsDeleted() (deleted bool) {
	return p.DeletedAt.Valid && p.DeletedBy.Valid
}


func (product Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(product.ToResponseFormat())
}


func (p Product) NewFromRequestFormat(req ProductRequestFormat) (newProduct Product, err error) {
	productId, _ := uuid.NewV4()
	newProduct = Product{
		Id:         productId,
		BrandId: 	req.BrandId,
		Name:       req.Name,
		CreatedAt:  time.Now(),
		CreatedBy:  req.UserId,
		UpdatedAt: 	null.TimeFrom(time.Now()),
		UpdatedBy: 	nuuid.From(req.UserId),
	}

	variants := make([]Variant, 0)
	for _, requestVariants := range req.Variants {
		variant := Variant{}
		variant = variant.NewFromRequestFormat(requestVariants, productId, req.UserId)
		variants = append(variants, variant)
	}

	newProduct.Variants = variants

	return
}



func (p *Product) SoftDelete(userID uuid.UUID) (err error) {
	if p.IsDeleted() {
		return failure.Conflict("softDelete", "Product", "already marked as deleted")
	}

	p.DeletedAt = null.TimeFrom(time.Now())
	p.DeletedBy = nuuid.From(userID)

	return
}


func (p Product) ToResponseFormat() ProductResponseFormat {
	resp := ProductResponseFormat{
		Id: 			p.Id,
		BrandId: 		p.BrandId,
		Name: 			p.Name,
		CreatedAt:      p.CreatedAt,
		CreatedBy:     	p.CreatedBy,
		UpdatedAt:      p.UpdatedAt,
		UpdatedBy:     	p.UpdatedBy.Ptr(),
		DeletedAt:      p.DeletedAt,
		DeletedBy:     	p.DeletedBy.Ptr(),
		Variants: 		make([]VariantResponseFormat, 0),	
	}

	for _, variant := range p.Variants {
		resp.Variants = append(resp.Variants, variant.ToResponseFormat())
	}

	return resp
}


func (p *Product) Update(req ProductRequestFormat, userID uuid.UUID) (err error) {
	variants := make([]Variant, 0)
	for _, requestVariant := range req.Variants {
		variant := Variant{}
		variant = variant.NewFromRequestFormat(requestVariant, p.Id, userID)
		variants = append(variants, variant)
	}

	p.Variants = variants
	p.BrandId = req.BrandId
	p.Name = req.Name
	p.UpdatedAt = null.TimeFrom(time.Now())
	p.UpdatedBy = nuuid.From(userID)

	return
}


type ProductRequestFormat struct {
	UserId 			uuid.UUID				`json:"userId" validate:"required"`
	BrandId 		uuid.UUID				`json:"brandId" validate:"required"`
	Name			string	 				`json:"name" validate:"required"`
	Variants       []VariantRequestFormat 	`json:"variants" validate:"required,dive,required"`
}


type ProductResponseFormat struct {
	Id   			uuid.UUID				`json:"id"`
	BrandId 		uuid.UUID				`json:"name"`
	Name			string	 				`json:"type"`
	Variants        []VariantResponseFormat `json:"variants"`
	CreatedAt       time.Time  				`json:"createdAt" `
	CreatedBy     	uuid.UUID   			`json:"createdBy"`
	UpdatedAt       null.Time   			`json:"updatedAt,omitempty"`
	UpdatedBy     	*uuid.UUID				`json:"updatedBy,omitempty"`
	DeletedAt       null.Time   			`json:"deletedAt,omitempty"`
	DeletedBy     	*uuid.UUID				`json:"deletedBy,omitempty"`
}


type Variant struct {
	Id          	uuid.UUID 		`db:"id" validate:"required"`
	ProductId       uuid.UUID 		`db:"productId" validate:"required"`
	Name         	string    		`db:"name" validate:"required"`
	Price 			float64    		`db:"price" validate:"required"`
	CreatedAt       time.Time  		`db:"createdAt" `
	CreatedBy     	uuid.UUID   	`db:"createdBy"`
	UpdatedAt       null.Time   	`db:"updatedAt"`
	UpdatedBy     	nuuid.NUUID		`db:"updatedBy"`
	DeletedAt       null.Time   	`db:"deletedAt"`
	DeletedBy     	nuuid.NUUID		`db:"deletedBy"`
}

func (v Variant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.ToResponseFormat())
}

func (v Variant) NewFromRequestFormat(req VariantRequestFormat, productID uuid.UUID, userID uuid.UUID) (variant Variant) {
	variantID, _ := uuid.NewV4()
	variant = Variant{
		Id: 		variantID,
		ProductId: 	productID,
		Name: 		req.Name,
		Price: 		req.Price,
		CreatedAt:  time.Now(),
		CreatedBy:  userID,
		UpdatedAt: 	null.TimeFrom(time.Now()),
		UpdatedBy: 	nuuid.From(userID),
	}
	return
}



func (v *Variant) ToResponseFormat() VariantResponseFormat {
	return VariantResponseFormat{
		Id: 		v.Id,
		ProductId: 	v.ProductId,
		Name: 		v.Name,
		Price: 		v.Price,
		CreatedAt: 	v.CreatedAt,
		CreatedBy: 	v.CreatedBy,
		UpdatedAt: 	v.UpdatedAt,
		UpdatedBy: 	v.UpdatedBy,
		DeletedAt: 	v.DeletedAt,
		DeletedBy: 	v.DeletedBy,
	}
}


type VariantRequestFormat struct {
	Name 		string		`json:"name" validate:"required"`
	Price		float64		`json:"price" validate:"required"`
}


type VariantResponseFormat struct {
	Id          	uuid.UUID 		`json:"id"`
	ProductId       uuid.UUID 		`json:"productId"`
	Name         	string    		`json:"name"`
	Price 			float64    		`json:"price"`
	CreatedAt       time.Time  		`json:"createdAt" `
	CreatedBy     	uuid.UUID   	`json:"createdBy"`
	UpdatedAt       null.Time   	`json:"updatedAt,omitempty"`
	UpdatedBy     	nuuid.NUUID		`json:"updatedBy,omitempty"`
	DeletedAt       null.Time   	`json:"deletedAt,omitempty"`
	DeletedBy     	nuuid.NUUID		`json:"deletedBy,omitempty"`
}




