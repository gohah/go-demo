package model

import (
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/jmoiron/sqlx"
)

type ProductModel struct {
	
}

type Product struct {
	ProductId   int    `db:"id"`
	ProductName string `db:"name"`
	Total       int    `db:"total"`
	Status      int    `db:"status"`
}

func NewProductModel() *ProductModel {
	productModel := &ProductModel{
	}

	return productModel
}

func (p *ProductModel) GetProductList() (list []*Product, err error) {
	sql := "select id, name, total, status from product"
	err = Db.Select(&list, sql)
	if err != nil {
		logs.Warn("select from mysql failed, err:%v sql:%v", err, sql)
		return
	}

	return
}


func (p *ProductModel) CreateProduct(product *Product) (err error) {
	
	sql := "insert into product(name, total, status)values(?,?,?)"
	_, err = Db.Exec(sql, product.ProductName, product.Total, product.Status)
	if err != nil {
		logs.Warn("select from mysql failed, err:%v sql:%v", err, sql)
		return
	}
	
	logs.Debug("insert into database succ")
	return
}
