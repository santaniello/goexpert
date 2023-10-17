package main

import (
	"database/sql"
	"fmt"

	/*
		 O _ faz com que consigamos compilar nosso programa visto que náo estamos usando esse pacote diretamente
		pois ele é o driver do mysql qu é usado pelo GO para se conectar ao banco. quem usa no caso, é a funcao sql.Open
	*/
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	product := NewProduct("Notebook", 18000)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	product.Name = "Notebook LG"
	fmt.Println("Testando alteracao de produto ")
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	fmt.Println("Tenstando buscar produto com o id : " + product.ID)
	productFounded, _ := findOneProduct(db, product.ID)
	fmt.Println(productFounded)

	allProducts, err := findAllProducts(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listando todos os valores")
	for _, v := range allProducts {
		fmt.Println(v)
	}

	fmt.Println("Deletando produto com o id " + product.ID)
	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
}

/*
Por que estamos passando um ponteiro de Product e náo estamos fazendo a passagem por valor ?

No Go, Quando passamos uma variável por valor para uma função em Go (ou em muitas outras linguagens, na verdade), o que acontece é que uma cópia dessa variável é feita. Então, se a variável for uma estrutura grande, toda essa estrutura será copiada, o que pode ser custoso em termos de tempo e uso de memória.

Suponha que temos uma estrutura BigStruct que ocupa, digamos, 1 MB de memória. Se passarmos uma variável do tipo BigStruct por valor para uma função, 1 MB de memória será copiado. Se fizermos isso muitas vezes, rapidamente estaremos falando de uma quantidade significativa de dados sendo copiada de um lugar para outro na memória.

Agora, se passarmos um ponteiro para essa estrutura, estamos essencialmente passando apenas um endereço de memória (que é o que um ponteiro é). Isso é muito mais leve - normalmente um ponteiro ocupa apenas 4 bytes (para sistemas de 32 bits) ou 8 bytes (para sistemas de 64 bits), independentemente do tamanho da estrutura para a qual ele aponta.

Então, em nosso exemplo anterior, em vez de copiar 1 MB de dados, estaríamos copiando apenas 4 ou 8 bytes.

A passagem por valor é ideal para structs pequenas e quando queremos garantir a imutabilidade da estrutura que estamos passando.

No entanto, é bom mencionar que, para estruturas pequenas, a diferença de desempenho pode ser negligenciável e, às vezes, até mesmo contraproducente usar ponteiros devido à indireção (acessar dados através de um ponteiro requer um passo adicional de desreferência). Portanto, a decisão de usar ponteiros versus passagem por valor geralmente depende do tamanho da estrutura e dos requisitos específicos do programa.
*/
func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products(id, name, price) values (?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set  name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func findOneProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select * from products  where id = ?")
	defer stmt.Close()
	var p Product
	// Na linha de baixo estamos selecionando apenas 1 registro da tabela e atribuindo os valores para o ponteiro p
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}

	return &p, nil

}

func findAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products  where id = ?")
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
