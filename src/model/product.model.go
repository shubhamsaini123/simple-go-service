package models
type ProductModel
import (
	"entities"
)
type ProductModel struct{

}
func (productModel *ProductModel) FindAll()([]entites.product,error){
	db,error :=config.GETDB()
	if err !=nil{
		return err, nil
	}else{
		var products []entities.Product
		db.Find(&products)
		return products, nil

	}
}