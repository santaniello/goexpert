package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	/*
		Quando queremos representar um relacionamento usando gorm,
		precisamos ter o id da entidade e o modelo da mesma
	*/
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

/*
Diferente da categoria onde um produto pode pertencer a uma categoria (belongs_to), aqui um produto tem um numero serial
e somente esse produto pode conter esse numero serial.
*/
type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// Criando as tabelas de produto e categoria e serialnumber
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// criando a categoria
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	// Criando o produto passando o id da categoria que acabou de ser criada
	product := Product{Name: "Notebook", Price: 1000, CategoryID: category.ID}
	db.Create(&product)

	// create serial number
	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: product.ID,
	})

	var products []Product
	/*
		O Preload faz com que ao buscarmos todos os produtos do banco de dados através da query Find,
		todas as entidades categorias vinculadas a cada produto seja trazida do banco junto com o produto. Se náo fizessemos isso,
		o objeto Category seria nil e só traria o CategoryID
	*/
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, p := range products {
		fmt.Println(p.Name, p.Category.Name, p.SerialNumber.Number)
	}

}
