# Contextos

[Documentação](https://pkg.go.dev/context)

## O que são contextos ?

O pacote context define a interface Context, que transporta prazos, cancelamentos, e outros valores
entre os processos. Uma instância de Context é imutável; ou seja, para modificar um contexto, você 
cria uma nova instância dele.

O pacote context em Go é uma ferramenta poderosa para gerenciar cancelamentos, prazos e transporte 
de valores entre goroutines e funções. Ele ajuda a escrever programas mais robustos e resilientes, 
especialmente em situações onde o controle sobre a execução concorrente é necessário.

## Para que Servem ?

1. **Cancelamentos (WithCancel):** Imagine que você tenha um serviço que faz várias chamadas a outros serviços. Se um dos serviços demorar muito, talvez você queira cancelar todas as outras chamadas. O contexto permite que você propague um sinal de cancelamento para todas as goroutines que foram derivadas dele.

2. **Timeouts e Deadlines (WithDeadline ou WithTimeout):** Junto com cancelamentos, você pode especificar um prazo (deadline) após o qual o contexto será cancelado. Isso é útil para definir um tempo máximo que um processo ou tarefa pode levar.

3. **Transporte de Valores (WithValue):** O contexto permite transportar valores específicos de um processo para outro. No entanto, é importante ter cautela ao usar esta funcionalidade, pois pode tornar o código difícil de entender e manter.

## Como usar ?

Aqui estão os principais métodos e funções que você verá ao trabalhar com contextos:

- **context.Background():** Retorna um contexto vazio. Geralmente é o ponto de partida para a criação de novos contextos.

- **ctx.WithCancel(parentCtx):** Retorna um novo contexto e uma função cancel que pode ser usada para cancelar o contexto.

- **ctx.WithDeadline(parentCtx, deadline):** Retorna um novo contexto que será cancelado no prazo especificado.

- **ctx.WithTimeout(parentCtx, timeout):** Retorna um novo contexto que será cancelado após o timeout especificado.

- **ctx.WithValue(parentCtx, key, val):** Retorna um novo contexto com o valor associado à chave.

## Select 

O select é uma estrutura de controle em Go (também conhecido como Golang) usada principalmente 
para lidar com múltiplos canais. Ela permite que o programa espere até que uma das comunicações 
possa seguir em frente, o que é útil para lidar com operações concorrentes.

Exemplo: 

```go
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    log.Println("Request iniciada")
    defer log.Println("Request finalizada")
    select {
        case <-time.After(5 * time.Second):
            log.Println("Request processada com sucesso")
            w.Write([]byte("Request processada com sucesso"))
        case <-ctx.Done():
            log.Println("Request cancelada pelo cliente")
    }
}
```

No exemplo acima, o select que é muito similar a um switch case pois ele esta analisando dois processamentos que podem ocorrer 
a qualquer momento. No primeiro case, ele irá aguardar 5 segundos para entrar nesse case (time.After(5 * time.Second))
e caso esses 5 segundos sejam atingisos, a mensagem "Request processada com sucesso" será exibida.

Em paralelo a isso, o usuário pode cancelar esse processamento antes dos 5 segundos apertando um ctrl + c e sendo assim, vai cair no segundo case pois a função ctx.Done()
verifica se aquele contexto foi cancelado e caso isso ocorra, a mensagem "Request cancelada pelo cliente" será exibida.

## Qual a diferença de um switch case para um select ?

Tanto o select quanto o switch são estruturas de controle de fluxo em Go (Golang), mas eles têm usos e comportamentos diferentes.

1. select:
    - É usado exclusivamente para operações com canais (channels).
    - Permite que o programa espere em múltiplos canais até que um deles possa proceder com uma operação de envio ou recebimento.
    - Não possui uma condição de "default" obrigatória, mas pode incluir uma para tratar situações onde nenhuma das outras operações de canal está pronta.
    - Se vários canais estiverem prontos ao mesmo tempo, o select escolhe um caso aleatoriamente para executar.
2. switch :
    - É uma estrutura de controle geral usada para testar valores e tipos.
    - Avalia uma expressão ou valor e compara com possíveis casos.
    - Como o select, o switch pode ter uma cláusula "default", mas ela é opcional.
    - Além do switch tradicional baseado em valores, Go também suporta o switch de tipo, que pode ser usado para descobrir o tipo de uma interface.

Em resumo, enquanto o switch é uma ferramenta geral para testar valores e tipos, o select é especializado em lidar com operações de canal em cenários de concorrência. Ambos são essenciais para a programação em Go, especialmente quando se lida com goroutines e concorrência.

## O que é o operador <-

Em Go (Golang), o símbolo <- é utilizado principalmente no contexto de operações com canais (channels). Canais são uma ferramenta poderosa em Go para comunicação entre goroutines, permitindo a troca segura de dados entre elas.

O símbolo <- é usado tanto para enviar quanto para receber valores de um canal. 


