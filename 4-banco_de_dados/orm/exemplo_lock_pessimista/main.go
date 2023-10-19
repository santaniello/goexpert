package main

import (
	"gorm.io/gorm/clause"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})
	/*
		Isso inicia uma nova transação no banco de dados. Transações são úteis para garantir que uma série de operações sejam executadas de forma atômica (ou todas são bem-sucedidas ou nenhuma é).
	*/
	tx := db.Begin()
	var c Category
	/*
	  Abaixo, O código tenta bloquear e obter a primeira entrada (com ID=1) da tabela de Category para atualização. O bloqueio garante que nenhuma outra transação possa modificar esta entrada até que a atual transação seja concluída.
	*/
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}

	c.Name = "Eletronicos2"
	tx.Debug().Save(&c)
	/*
		Isso confirma (ou "commita") a transação, efetivamente salvando todas as alterações feitas desde o início da transação no banco de dados e liberando o bloqueio.
	*/
	tx.Commit()
}
