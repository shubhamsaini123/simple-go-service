package entities
type Product struct{
	Id 		int
	Name 	string
	Price   float64
	Quantity int
	Status  bool
	Created time.Time
}
func (product *Product) TableName() string{
	return "product"
}
func(product Product) ToString() string{
	return fmt.Sprintf("id:%d\nname:%s\nprice:%0.1f\nquantity:%d\nstatus:%t\ncreated:%s",
    product.Id,product.Name,product.Price,product.Quantity,product.Status,product.Created.Format("2019-09-20"))
}