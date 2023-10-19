package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

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
	// Criando as tabelas de produto e categoria
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// criando a categoria
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	// Criando o produto passando o id da categoria que acabou de ser criada
	product := Product{Name: "Notebook", Price: 1000, CategoryID: category.ID}
	db.Create(&product)

	// // create serial number
	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: product.ID,
	})

	var categories []Category
	/*
		 Model especifica o modelo que queremos rodar nossas operacoes de consulta ... nesse caso category e carregamos todos os produtos vinvulados a essas categorias
		 .Error: Se houver um erro na consulta, ele é atribuído à variável err.

		O Preload("Products.SerialNumber") foi feito pois queremos saber qual o serial number esta atrelado a aquele produto e ele traz tanto os produtos quanto os serial numbers atrelados a esses produtos, ou seja, nao preciso ter um Preload("Products").Preload("Products.SerialNumber")
		Se deixassemos apenas ("Products"), ao tentar imprimir o serial number nós teriamos um em branco como resposta ! isso devido ao fato de o mesmo nao ter sido carregado junto com o Product atrelado a aquela categoria...
		Se adicionassemos mais um Preload("SerialNumber"), teriamos um erro de SQL pois ele iria tentar procurar essa estrutura no modelo Category e nao encontraria !

	*/
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			println("- ", product.Name, "Serial Number:", product.SerialNumber.Number)
		}
	}
}
