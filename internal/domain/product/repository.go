package product

import (
	"database/sql"
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
		selectProduct				 string
		selectProductStatusSorted	 string
		selectVariant				 string
		insertProduct                string
		insertVariantBulk            string
		insertVariantBulkPlaceholder string
		updateProduct				 string	
		selectImages				 string
	}{	selectProduct: `
		SELECT
			product.id,
			product.brandId,
			product.name,
			product.stock,
			product.createdAt,
			product.createdBy,
			product.updatedAt,
			product.updatedBy,
			product.deletedAt,
			product.deletedBy
		FROM product 
		`,
		selectProductStatusSorted: `
		SELECT 
			b.name AS brandName,
			p.name AS productName,
			v.name AS variantName,
			v.id  AS variantId,
			v.price AS price,
			v.stock AS stock,
			CASE 
				WHEN p.stock > 0 THEN 'Ready'
				ELSE 'Out of Stock'
			END AS status,
			p.updatedBy AS updatedBy
		FROM 
			Brand b
		JOIN 
			Product p ON b.id = p.brandId
		JOIN 
			Variant v ON p.id = v.productId
		`,
		selectVariant: `
		SELECT
			id,
			productId,
			name,
			price,
			stock,
			createdAt,
			createdBy,
			updatedAt,
			updatedBy,
			deletedAt,
			deletedBy
		FROM variant
		`,
		insertProduct: `
			INSERT INTO product (
				id,
				brandId,
				name,
				stock,
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
				:stock,
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
				stock,
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
			:stock,
			:createdAt,
			:createdBy,
			:updatedAt,
			:updatedBy,
			:deletedAt,
			:deletedBy)`,
		
		updateProduct: `
			UPDATE product
			SET
				brandId = :brandId,
				name = :name,
				stock = :stock,
				createdAt = :createdAt,
				createdBy = :createdBy,
				updatedAt = :updatedAt,
				updatedBy = :updatedBy,
				deletedAt = :deletedAt,
				deletedBy = :deletedBy
			WHERE id = :id 
		`,

		selectImages: `
		SELECT url FROM image
		`,
	}
)


type ProductRepository interface {
	Create(product Product) (err error)
	Update(product Product) (err error)
	HardDelete(id uuid.UUID) (err error)
	ResolveByID(id uuid.UUID) (product Product, err error)
	ExistsByID(id uuid.UUID) (exists bool, err error)
	ExistsByUserID(userId uuid.UUID) (exists bool, err error)
	ResolveVariantsByProductIDs(ids []uuid.UUID) (variants []Variant, err error)
	ResolveImagesByVariantID(ids uuid.UUID) (images []string, err error)
	ReadPagination(limit int, offset int) (products []Product, err error)
	ReadStatusSorted() (products []ProductStatus, err error)
	ReadByBrandName(brandName string) (products []ProductStatus, err error)
	ReadByProductName(productName string) (products []ProductStatus, err error)
	ReadByVariantName(variantName string) (products []ProductStatus, err error)
	ReadByStatus(status string) (products []ProductStatus, err error)
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


func (r *ProductRepositoryMySQL) ExistsByUserID(id uuid.UUID) (exists bool, err error) {
	err = r.DB.Read.Get(
		&exists,
		"SELECT COUNT(id) FROM user WHERE user.id = ?",
		id.String())
	if err != nil {
		logger.ErrorWithStack(err)
	}

	return
}



func (r *ProductRepositoryMySQL) ResolveByID(id uuid.UUID) (product Product, err error) {
	err = r.DB.Read.Get(
		&product,
		productQueries.selectProduct+" WHERE product.id = ?",
		id.String())
	if err != nil && err == sql.ErrNoRows {
		err = failure.NotFound("product")
		logger.ErrorWithStack(err)
		return
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


func (r *ProductRepositoryMySQL) Update(product Product) (err error) {
	exists, err := r.ExistsByID(product.Id)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}

	if !exists {
		err = failure.NotFound("product")
		logger.ErrorWithStack(err)
		return
	}

	// transactionally update the product
	// strategy:
	// 1. delete all the product's items
	// 2. create a new set of product's items
	// 3. update the product
	return r.DB.WithTransaction(func(tx *sqlx.Tx, e chan error) {
		if err := r.txDeleteVariants(tx, product.Id); err != nil {
			e <- err
			return
		}

		if err := r.txCreateVariants(tx, product.Variants); err != nil {
			e <- err
			return
		}

		if err := r.txUpdate(tx, product); err != nil {
			e <- err
			return
		}

		e <- nil
	})
}

func (r *ProductRepositoryMySQL) ReadPagination(limit int, page int) (products []Product, err error) {
	err = r.DB.Read.Select(
		&products,
		productQueries.selectProduct+" 	LIMIT ? OFFSET ?", limit, page*limit)
	if err != nil && err == sql.ErrNoRows {
		err = failure.NotFound("products")
		logger.ErrorWithStack(err)
		return
	}
	return
}

func (r *ProductRepositoryMySQL) ReadStatusSorted() (products []ProductStatus, err error) {
	err = r.DB.Read.Select(
		&products,
		productQueries.selectProductStatusSorted+" ORDER BY p.createdAt DESC,v.stock DESC; ")
	if err != nil && err == sql.ErrNoRows {
		err = failure.NotFound("products")
		logger.ErrorWithStack(err)
		return
	}
	return
}

func (r *ProductRepositoryMySQL) ReadByBrandName(brandName string) (products []ProductStatus, err error) {
	err = r.DB.Read.Select(
		&products,
		"WITH productDetails AS(" +productQueries.selectProductStatusSorted+ ") SELECT * FROM productDetails WHERE brandName = ? ORDER BY stock DESC;", brandName)
	if err != nil && err == sql.ErrNoRows {
		err = failure.NotFound("products")
		logger.ErrorWithStack(err)
		return
	}
	return
}

func (r *ProductRepositoryMySQL) ReadByProductName(productName string) (products []ProductStatus, err error) {
	err = r.DB.Read.Select(
		&products,
		"WITH productDetails AS(" +productQueries.selectProductStatusSorted+ ") SELECT * FROM productDetails WHERE productName = ? ORDER BY stock DESC;", productName)
	if err != nil && err == sql.ErrNoRows {
		err = failure.NotFound("products")
		logger.ErrorWithStack(err)
		return
	}
	return
}

func (r *ProductRepositoryMySQL) ReadByVariantName(variantName string) (products []ProductStatus, err error) {
	err = r.DB.Read.Select(
		&products,
		"WITH productDetails AS(" +productQueries.selectProductStatusSorted+ ") SELECT * FROM productDetails WHERE variantName = ? ORDER BY stock DESC;", variantName)
	if err != nil && err == sql.ErrNoRows {
		err = failure.NotFound("products")
		logger.ErrorWithStack(err)
		return
	}
	return
}

func (r *ProductRepositoryMySQL) ReadByStatus(status string) (products []ProductStatus, err error) {
	err = r.DB.Read.Select(
		&products,
		"WITH productDetails AS(" +productQueries.selectProductStatusSorted+ ") SELECT * FROM productDetails WHERE status = ? ORDER BY stock DESC;", status)
	if err != nil && err == sql.ErrNoRows {
		err = failure.NotFound("products")
		logger.ErrorWithStack(err)
		return
	}
	return
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
			"stock":		v.Stock,
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


func (r *ProductRepositoryMySQL) txDeleteVariants(tx *sqlx.Tx, productId uuid.UUID) (err error) {
	_, err = tx.Exec("DELETE FROM variant WHERE productId  = ?", productId.String())
	return
}


func (r *ProductRepositoryMySQL) txUpdate(tx *sqlx.Tx, product Product) (err error) {
	stmt, err := tx.PrepareNamed(productQueries.updateProduct)
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



func (r *ProductRepositoryMySQL) ResolveVariantsByProductIDs(ids []uuid.UUID) (variants []Variant, err error) {
	if len(ids) == 0 {
		return
	}

	query, args, err := sqlx.In(productQueries.selectVariant+" WHERE variant.productId IN (?)", ids)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}

	err = r.DB.Read.Select(&variants, query, args...)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}

	return
}



func (r *ProductRepositoryMySQL) ResolveImagesByVariantID(ids uuid.UUID) (images []string, err error) {
	if len(ids) == 0 {
		return
	}

	query, args, err := sqlx.In(productQueries.selectImages+" WHERE image.variantId = ?", ids)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}

	err = r.DB.Read.Select(&images, query, args...)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}

	return
}

func (r *ProductRepositoryMySQL) HardDelete(id uuid.UUID) (err error) {
	exists, err := r.ExistsByID(id)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}

	if !exists {
		err = failure.NotFound("product")
		logger.ErrorWithStack(err)
		return
	}


	return r.DB.WithTransaction(func(tx *sqlx.Tx, e chan error) {
		if err := r.txDeleteVariants(tx, id); err != nil {
			e <- err
			return
		}

		if err := r.txDeleteProduct(tx, id); err != nil {
			e <- err
			return
		}


		e <- nil
	})
}


func (r *ProductRepositoryMySQL) txDeleteProduct(tx *sqlx.Tx, productId uuid.UUID) (err error) {
	_, err = tx.Exec("DELETE FROM product WHERE id  = ?", productId.String())
	return
}
