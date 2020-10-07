
package main
import (
	"models"
    "fmt"
)
func main(){
   Demo1()
}
func Demo1(){
  var productModel models.ProductModel
  products,_ := productModel.FindALL()
  for _,product:=range products{
	  fmt.Println(product.ToString())
	  fmt.Println(-------------------)
  }
}