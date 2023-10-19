# Banco de Dados

## Lock Otimista VS Pessimista

### Lock Otimista

O "Lock Otimista" (ou "Optimistic Locking", em inglês) é uma estratégia usada para gerenciar o acesso concorrente a um recurso de dados sem a utilização de bloqueios tradicionais durante a leitura do recurso. Em vez de bloquear um recurso durante todo o período de leitura e atualização, o controle de concorrência otimista permite que vários processos ou threads leiam o recurso simultaneamente, contando com o fato de que conflitos (duas ou mais operações tentando atualizar o mesmo recurso ao mesmo tempo) são relativamente raros.

O Lock Otimista usa basicamente o versionamento do registro para saber se o mesmo foi alterado ou não. Esse versionamento pode ser atraves de timestamp, numero de versão e etc ...

Como funciona ?

Basicamente, um processo le um registro o qual pretende-se fazer uma alteracao e nele esta contido algum tipo de marca que indica a versao daquele registro (numero de versao, timestamp e etc...).
Apos isso, o registro e alterado e antes de o processo escrever suas modificacoes na tabela, ele verifica se essa marca foi alterada. 
Se a marca nao foi alterada, as alteracoes sao escritas no banco pois significa que ninguem alterou esse registro.
Se a marca foi alterada, significa que algum outro processo alterou o registro desde a leitura inicial e portanto, temos um conflito.

***Resolução de Conflitos:*** Se ocorrer um conflito (a tentativa de atualização é rejeitada), o processo ou thread tem várias opções: 
- Pode abandonar suas mudanças;
- Tentar a operação novamente baseado em um numero de retentativas;
- Mesclar suas mudanças com as mudanças concorrentes; 
- Notificar o usuário ou administrador;

### Exemplo de Lock Otimista

Imagine que abaixo temos uma tabela de banco de dados:


name      | email | versao
--------- |-------| --------
Felipe    | f@f   | 1
Carlos    | c@c   |  2

Repare que a coluna versao seria a nossa marca e ela quem faz o versionamento do registro.

Vantagens:

- Reduz a contenção, já que os bloqueios não são mantidos durante toda a duração da operação.
- É especialmente útil em ambientes onde os conflitos são raros, pois evita o overhead de bloqueios constantes.

Desvantagens:

- Pode haver um overhead adicional se os conflitos forem frequentes, pois as operações precisam ser refeitas.
- A resolução de conflitos pode ser complexa em alguns cenários.

O bloqueio otimista é frequentemente implementado em sistemas de gerenciamento de banco de dados, ORM (Object-Relational Mapping) e outros sistemas de armazenamento de dados.

### Lock Pessimista

O "Lock Pessimista" (ou "Pessimistic Locking", em inglês) é uma estratégia usada para gerenciar o acesso concorrente a um recurso de dados garantindo que, uma vez que um processo ou thread tenha bloqueado esse recurso para leitura ou atualização, nenhum outro processo ou thread possa acessar o recurso até que o bloqueio seja liberado.

O Lock Pessimista bloqueia basicamente "Trava" um registro diretamente no banco de dados não permitindo que o mesmo seja alterado por outro processo.

Como Funciona ?

Quando um processo ou thread deseja ler ou modificar um recurso, ele solicita um bloqueio.
Se o recurso não estiver bloqueado por outra transação, o sistema de gerenciamento de dados concede o bloqueio, evitando que qualquer outro processo ou thread acesse o recurso até que o bloqueio seja liberado.
Com o recurso bloqueado, o processo ou thread pode proceder para ler ou modificar o recurso.
Uma vez que as operações no recurso são concluídas, o bloqueio é liberado, tornando o recurso disponível para outros processos ou threads.

#### Exemplo de Lock Pessismista 

```go
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Category struct {
	ID   int
	Name string
}

func main() {
	// Conectar ao banco de dados
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Iniciar transação
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Executar o bloqueio pessimista
	var c Category
	err = tx.QueryRow("SELECT id, name FROM categories WHERE id = ? FOR UPDATE", 1).Scan(&c.ID, &c.Name)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// Aqui você pode fazer alterações na linha bloqueada
	c.Name = "Eletronicos2"
	_, err = tx.Exec("UPDATE categories SET name = ? WHERE id = ?", c.Name, c.ID)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// Commitar as alterações
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Categoria atualizada com sucesso!")
}

```

Nesse projeto, temos um exemplo de bloquei pessimista usando o ORM GORM, esta na pasta orm/exemplo_lock_pessimista

Vantagens:

- Fornece uma garantia direta de que, uma vez que o recurso é bloqueado, nenhum outro processo ou thread pode interferir até que o bloqueio seja liberado.
- Evita a necessidade de lidar com conflitos durante a fase de atualização, pois os conflitos são gerenciados na fase de aquisição de bloqueio.

- Desvantagens:

- Pode levar a uma contenção significativa, especialmente em sistemas com alto grau de concorrência, pois os recursos ficam indisponíveis para outros processos ou threads por mais tempo.
- Pode levar a "deadlocks" ou impasses, situações em que dois ou mais processos ou threads estão esperando uns pelos outros para liberar recursos, resultando em um bloqueio permanente.

Enquanto o bloqueio otimista assume que conflitos são raros e lida com eles quando ocorrem na fase de atualização, o bloqueio pessimista assume que conflitos são prováveis e proativamente bloqueia recursos para evitar interferência.

O bloqueio pessimista é frequentemente usado em sistemas que requerem garantias estritas de isolamento, onde a integridade dos dados é crítica, ou em cenários onde a probabilidade de conflito é alta. É comum em muitos sistemas de gerenciamento de banco de dados relacionais tradicionais.

