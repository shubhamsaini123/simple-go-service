package config
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
  )
  
  func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
   if err !=nil{
	   return nil,err
   }  
   return db, nil
}