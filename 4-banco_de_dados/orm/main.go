package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	criarTabelaProduct(db)

	//criarProduto(db, &Product{
	//	Name:  "Notebook",
	//	Price: 1000,
	//})
	//
	//products := []Product{
	//	{Name: "Notebook", Price: 2000},
	//	{Name: "Mouse", Price: 20},
	//	{Name: "Keyboard", Price: 100},
	//}
	//
	//criarVariosProdutos(db, &products)

	//fmt.Println("Resultado da Query : findOneById")
	//fmt.Println(findOneById(db, 3))
	//fmt.Println("Resultado da Query : findOneByName")
	//fmt.Println(findOneByName(db, "Mouse"))
	//fmt.Println("Resultado da Query : findAllProducts")
	//for _, v := range *findAllProducts(db) {
	//	fmt.Println(v)
	//}
	//fmt.Println("Resultado da Query : findAllProductsComPaginacao")
	//for _, v := range *findAllProductsComPaginacao(db) {
	//	fmt.Println(v)
	//}
	//fmt.Println("Resultado da Query : findComWhere")
	//for _, v := range *findComWhere(db) {
	//	fmt.Println(v)
	//}
	//fmt.Println("Resultado da Query : findComLike")
	//for _, v := range *findComLike(db) {
	//	fmt.Println(v)
	//}
	//fmt.Println("Resultado da Query : updateProduct")
	//updateProduct(db, 1)
	//fmt.Println(findOneById(db, 1))
	//
	//fmt.Println("Resultado da Query : deleteProduct")
	//deleteProduct(db, 1)
	//fmt.Println(findOneById(db, 1))

}

func criarProduto(db *gorm.DB, product *Product) {
	db.Create(product)
}

func criarVariosProdutos(db *gorm.DB, product *[]Product) {
	db.Create(product)
}

/*
A funcao Auto Migrate cria uma tabela no banco de dados baseado na struct que estamos passando
*/
func criarTabelaProduct(db *gorm.DB) {
	db.AutoMigrate(&Product{})
}

// ******** Consultas ********

func findOneById(db *gorm.DB, primaryKeyOrder int) *Product {
	var product Product
	db.First(&product, primaryKeyOrder)
	return &product
}

func findOneByName(db *gorm.DB, name string) *Product {
	var product Product
	db.First(&product, "name = ?", name)
	return &product
}

func findAllProducts(db *gorm.DB) *[]Product {
	var products []Product
	db.Find(&products)
	return &products
}

func findAllProductsComPaginacao(db *gorm.DB) *[]Product {
	var products []Product
	db.Limit(2).Offset(2).Find(&products)
	return &products
}

func findComWhere(db *gorm.DB) *[]Product {
	var products []Product
	db.Where("price > ?", 200).Find(&products)
	return &products
}

func findComLike(db *gorm.DB) *[]Product {
	var products []Product
	db.Where("name  LIKE ?", "%book%").Find(&products)
	return &products
}

func updateProduct(db *gorm.DB, id int) {
	product := findOneById(db, id)
	product.Name = "Produto Alterado"
	db.Save(&product)
}

func deleteProduct(db *gorm.DB, id int) {
	product := findOneById(db, id)
	db.Delete(&product)
}
