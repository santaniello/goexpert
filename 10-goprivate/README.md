# Trabalhando com repositórios privados no Go

## Variavel GOPRIVATE

A variável de ambiente GOPRIVATE é usada pelo Go (ou Golang) para especificar quais módulos são privados e, portanto, requerem autenticação para acessar. Quando você está trabalhando em um projeto Go que depende de módulos privados hospedados em um repositório Git, você pode usar a variável GOPRIVATE para informar ao Go que determinados domínios ou repositórios são privados e requerem autenticação para acessá-los.

A sintaxe básica para definir a variável GOPRIVATE é:

```bash
export GOPRIVATE=example.com
```

Para visualizar o valor da variavel GOPRIVATE, basta digitar o comando abaixo no terminal:

```bash
go env | grep PRIVATE
```

### Configurando a variavel GOPRIVATE

Foi criado ao projeto [fc-utils-private](https://github.com/santaniello/fc-utils-privat) no github para simular um repositório privado que queremos utilizar neste projeto. 

Oomando abaixo adiciona esepo a variavel GOPRIVATE:

```bash
 export GOPRIVATE=github.com/santaniello/fc-utils-private
```

**OBS: Se quisermos adicionar um outro repo, basta adicionar uma virgula e adicionar o outro repo.**

## Configurando credenciais de acesso via HTTPS (Login e Token)

Toda vez que o GO tentar baixar um repositório privado, ele tentará usar suas credencias e acesso para se logar no repositório remoto.

Para configurar essas credencias, precisamos alterar ou criar o seguinte arquivo:

```bash
 nano ~/.netrc
```

Nele, iremos inserir as seguintes infos:

```bash
machine github.com
login seu login do github
password seu token de acesso gerado no github
```

#### Como gerar o token de acesso no github ?

Vá até a seguinte pagina: https://github.com/settings/apps

Selecione Personal access tokens

e gere um classic token para sua app e é ele que você irá utilizar.

## Configurando credenciais de acesso via SSH 

Pata fazer a autenticação via ssh, iremos mexer no nosso arquivo .gitconfig geral (não o do projeto):

```bash
nano ~/.gitconfig
```

e irá adicionar a seguinte config:

```bash
[url "ssh://git@github.com/"]
        insteadOf = https://github.com/
```

## GO Proxy

O proxy Go, frequentemente referido como "Go Proxy", é um serviço oferecido pela comunidade Go para armazenar e servir módulos Go. Ele foi introduzido como parte do sistema de módulos Go a partir da versão 1.11 do Go.

Antes do Go Proxy, os desenvolvedores Go dependiam principalmente do repositório Git como fonte de módulos. Isso significava que, ao baixar módulos Go para seus projetos, o Go precisava clonar diretamente do repositório Git correspondente. No entanto, isso não era escalável e não fornecia uma solução robusta para versionamento e distribuição de módulos.

Com a introdução do Go Proxy, os desenvolvedores agora têm uma maneira centralizada e confiável de distribuir módulos Go. O Go Proxy armazena cópias de todos os módulos Go públicos, o que significa que os desenvolvedores podem baixar módulos diretamente do proxy, em vez de cloná-los diretamente do repositório Git original. Isso não apenas melhora significativamente a velocidade de download de módulos, mas também reduz a carga nos servidores do GitHub e de outros provedores de Git.

Além disso, o Go Proxy também resolve alguns problemas comuns relacionados à gestão de dependências em Go, como problemas de controle de versão e compatibilidade. Ele fornece uma maneira consistente e confiável de garantir que os projetos Go possam ser construídos e executados sem depender de alterações imprevistas nos repositórios de módulos.

Os desenvolvedores podem usar o Go Proxy padrão, que é o proxy oficial mantido pela equipe do Go, ou podem configurar seus próprios proxies para uso interno ou personalizado.

Site do GO Proxy: https://proxy.golang.org/

**OBS:** O Go Proxy seria como o Nexus para o Java.


### Baixando as dependências localmente

Para garantir que o build da nossa aplicação sempre ocorra independentemente se uma dependência sair fora do ar ou não, podemos utilizar o seguinte comando:

```bash
go mod vendor
```

Ele irá baixar todas dependências que nosso projeto utiliza para a pasta vendor (como se fosse uma node_modules).

**OBS:Esse recurso é bem pouc utilizado**



