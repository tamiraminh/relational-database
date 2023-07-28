package product

import (
	"fmt"
	"strings"

	"github.com/evermos/boilerplate-go/infras"
	"github.com/evermos/boilerplate-go/shared/failure"
	"github.com/evermos/boilerplate-go/shared/logger"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

var (
	productQueries = struct {
		insertProduct                string
		insertVariantBulk            string
		insertVariantBulkPlaceholder string
	}{
		insertProduct: `
			INSERT INTO product (
				id,
				brandId,
				name,
				createdAt,
				createdBy,
				updatedAt,
				updatedBy,
				deletedAt,
				deletedBy
			) VALUES (
				:id,
				:brandId,
				:name,
				:createdAt,
				:createdBy,
				:updatedAt,
				:updatedBy,
				:deletedAt,
				:deletedBy)`,

		insertVariantBulk: `
			INSERT INTO variant (
				id,
				productId,
				name,
				price,
				createdAt,
				createdBy,
				updatedAt,
				updatedBy,
				deletedAt,
				deletedBy
			) VALUES `,

		insertVariantBulkPlaceholder: `
			(:id,
			:productId,
			:name,
			:price,
			:createdAt,
			:createdBy,
			:updatedAt,
			:updatedBy,
			:deletedAt,
			:deletedBy)`,
	}
)


type ProductRepository interface {
	Create(product Product) (err error)
	// Update(product Product) (err error)
	// SoftDelete(id uuid.UUID) (err error)
	// HardDelete(id uuid.UUID) (err error)
	// ReadProduct() ( products []Product, err error)
	// ResolveByID(id uuid.UUID) (product Product, err error)
	ExistsByID(id uuid.UUID) (exists bool, err error)
	// ResolveItemsByProductIDs(ids []uuid.UUID) (fooItems []FooItem, err error)
}


type ProductRepositoryMySQL struct {
	DB *infras.MySQLConn
}


func ProvideProductRepositoryMySQL(db *infras.MySQLConn) *ProductRepositoryMySQL {
	s := new(ProductRepositoryMySQL)
	s.DB = db
	return s
}




func (r *ProductRepositoryMySQL) ExistsByID(id uuid.UUID) (exists bool, err error) {
	err = r.DB.Read.Get(
		&exists,
		"SELECT COUNT(id) FROM product WHERE product.id = ?",
		id.String())
	if err != nil {
		logger.ErrorWithStack(err)
	}

	return
}

func (r *ProductRepositoryMySQL) Create(product Product) (err error) {
	exists, err := r.ExistsByID(product.Id)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}

	if exists {
		err = failure.Conflict("create", "product", "already exists")
		logger.ErrorWithStack(err)
		return
	}

	return r.DB.WithTransaction(func(tx *sqlx.Tx, e chan error) {
		if err := r.txCreate(tx, product); err != nil {
			e <- err
			return
		}

		if err := r.txCreateVariants(tx, product.Variants); err != nil {
			e <- err
			return
		}

		e <- nil
	})
}




func (r *ProductRepositoryMySQL) txCreate(tx *sqlx.Tx, product Product) (err error) {
	stmt, err := tx.PrepareNamed(productQueries.insertProduct)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(product)
	if err != nil {
		logger.ErrorWithStack(err)
	}

	return
}


func (r *ProductRepositoryMySQL) txCreateVariants(tx *sqlx.Tx, variants []Variant) (err error) {
	if len(variants) == 0 {
		return
	}

	query, args, err := r.composeBulkInsertItemQuery(variants)
	if err != nil {
		return
	}

	stmt, err := tx.Preparex(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Stmt.Exec(args...)
	if err != nil {
		logger.ErrorWithStack(err)
	}

	return
}

func (r *ProductRepositoryMySQL) composeBulkInsertItemQuery(variants []Variant) (query string, params []interface{}, err error) {
	values := []string{}
	for _, v := range variants {
		param := map[string]interface{}{
			"id":    		v.Id,
			"productId":    v.ProductId,
			"name":         v.Name,
			"price": 		v.Price,
			"createdAt":    v.CreatedAt,
			"createdBy":   	v.CreatedBy,
			"updatedAt":  	v.UpdatedAt,
			"updatedBy":    v.UpdatedBy,
			"deletedAt":  	v.DeletedAt,
			"deletedBy":  	v.DeletedBy,
		}
		q, args, err := sqlx.Named(productQueries.insertVariantBulkPlaceholder, param)
		if err != nil {
			return query, params, err
		}
		values = append(values, q)
		params = append(params, args...)
	}
	query = fmt.Sprintf("%v %v", productQueries.insertVariantBulk, strings.Join(values, ","))
	return
}


