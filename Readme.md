# Validação, Verificação e Testes

Este projeto tem como objetivo servir de suporte ao curso de Verificação, Validação e Testes ministrado pelo professor [Daniel Fireman](mailto:danielfireman@gmail.com). Ele conterá todo o código fonte bem como um passo-a-passo da parte prática do curso.

Uma vez que será completamente efeutado pelo professor e o foco principal são os conceitos ministrados, o passo-a-passo poderia ser ministrado em qualquer linguagem de programação. Escolhemos [Go](http://golang.org) para esse curso. Para uma lista mais completa de motivos para a escolha de Go, por favor clique [aqui](#porque-go).



### Configuração do ambiente
* Instalar [Go runtime](http://golang.org)

Eu achei mais fácil usar gvm para instalação e manutenção de runtimes Go

```bash
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