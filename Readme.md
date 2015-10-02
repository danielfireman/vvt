# Validação, Verificação e Testes

Este projeto tem como objetivo servir de suporte ao curso de Verificação, Validação e Testes ministrado pelo professor [Daniel Fireman](mailto:danielfireman@gmail.com). Ele conterá todo o código fonte bem como um passo-a-passo da parte prática do curso.

Uma vez que será completamente efeutado pelo professor e o foco principal são os conceitos ministrados, o passo-a-passo poderia ser ministrado em qualquer linguagem de programação. Escolhemos [Go](http://golang.org) para esse curso. Para uma lista mais completa de motivos para a escolha de Go, por favor clique [aqui](#porque-go).

A aplicação a ser criada é um serviço de gerenciamento de TODOs. Para fins ilustrativos a aplicação armazena a lista de TODOs em memória. A API /todos tem 3 métodos: PUT, GET e DELETE; os quais são utilizados para adicionar, imprimir e remover items da lista. Existe também uma API GET /todolist, que imprime toda a lista de TODOs.

### Configuração do ambiente
Pessoas curiosas podem aprender o básico de [Go](http://golang.org) [aqui](https://tour.golang.org/welcome/1). Web app developers podem começar [aqui](https://golang.org/doc/articles/wiki/).

#### Instalação da runtime Go
Eu achei mais fácil usar [gvm](http://github.com/moovweb/gvm) para instalação e manutenção de runtimes Go. Intruções mais detalhadas [aqui](https://github.com/moovweb/gvm). É necessário ter  o comando  [curl](http://curl.haxx.se/) instalado.

```bash
$ bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
$ gvm install go1.5
$ gvm use go1.5 [--default]
$ echo "export GOPATH={SEU GOPATH}" >> ~/.bashrc
$ echo "export PATH=$PATH:$GOPATH/bin" >> ~/.bashrc
```

#### Dependências e código do servidor de TODOs
Para simplificar o código do servidor REST fazemos uso da biblioteca [labstack/echo](https://github.com/labstack/echo), outras bibliotecas também poderiam ser usadas, por exmplo [Gorilla Mux](https://github.com/gorilla/mux), [Gin](https://gin-gonic.github.io/gin/) e [Negroni](https://github.com/codegangsta/negroni). Esse passo só precisara ser executado uma vez.

    $ go get github.com/labstack/echo  # labstack/echo
    $ go get github.com/danielfireman/vvt

Após execução, o código fonte do servidor de TODOs poderá ser encontrado em :

    $GOPATH/src/github.com/danielfireman/vvt

### Executar servidor e utilizar o serviço
O comando abaixo compilará o código do servidor e executará o programa que aceitará requisções na porta 8999.

```bash
$ go run $GOPATH/src/github.com/danielfireman/vvt/cmd/server/main.go --port=8999
Server listening at localhost:8999
```

Pronto! A essa altura do campeonato temos o servidor executando e a API de gerênciamento de TODOs pronta para receber requisições na porta 8999. Vamos fazer alguns testes para entender melhor as funcionalidades do serviço. Utilizaremos [curl](http://curl.haxx.se/) para enviar as requisições.

```bash
 $ curl -H "Content-Type: application/json" -X POST -d '{"desc":"Comprar legumes"}' localhost:8999/todo
 {"Desc":"Comprar legumes"}
 $ curl -H "Content-Type: application/json" -X POST -d '{"desc":"Comprar verduras"}' localhost:8999/todo
 {"Desc":"Comprar verduras"}
```

Os comandos enviam requisições POST para http://localhost:8999/todo. A requisição é do tipo JSON e contém {"desc":"Comprar legumes"} e {"desc":"Comprar verduras"}. Ao final das execuções a lista de coisas a fazer terá dois items, como pode ser verificado através do comando abaixo:

```bash
 $ curl -H "Content-Type: application/json" -X GET localhost:8999/todolist
 ["Comprar legumes","Comprar verduras"]
```

Para remover o primeiro item da lista podemos usar o comando DELETE.

```bash
 $ curl -H "Content-Type: application/json" -X DELETE localhost:8999/todo/0
 $ curl -H "Content-Type: application/json" -X GET localhost:8999/todolist
 ["Comprar verduras"]
```

## Apendice
### Porque Go
Uma vez que será completamente efeutado pelo professor e o foco principal são os conceitos ministrados, o passo-a-passo poderia ser feito em qualquer linguagem de programação. Dentre os motivos que nos levaram a escolher [Go](http://golang.org) podemos listar:

* Sintaxe simples e enxuta, nos permitindo focar nos conceitos
* Suporte a testes de unidade embutido na linguagem: [pacote testing](https://golang.org/pkg/testing/)
    * Inclui benchmarks e perfis de processamento e memória
* Suporte primário a testes de integração embutidos na linguagem: [pacote httptest](https://golang.org/pkg/net/http/httptest/)
* Muitas ferramentas inspeção de código com simples instlação e utilização, por exemplo: 
    * Construções suspeitas: [Vet](https://golang.org/cmd/vet/) 
    * Errors de estilo: [GoLint](https://github.com/golang/lint)
    * Erros não verificados: [ErrCheck](http://github.com/kisielk/errcheck)
    * Injeção SQL: [SafeSQL](https://github.com/stripe/safesql) 
    * Expressões defer repetidas, structs com campos não-utilizados ou estruturada ineficientemente,  variavéis e constantes globaus não-utilizadas: [OpennotaCheck](https://github.com/opennota/check/)
* Debugger de simples instalação e utilização: [delve](https://github.com/derekparker/delve)
* Suportada por [drone.io](http://drone.io), ferramenta de entrega contínua
* Suportada por [Heroku](https://www.heroku.com/), plataforma onde será feito o deployment da aplicação