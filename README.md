# Go Projects Collection

## Overview

Este repositório reúne projetos de estudo em Go abordando: -
Concorrência com Goroutines e Channels - APIs REST com Gorilla Mux -
Testes unitários com package testing

------------------------------------------------------------------------

# Projeto 1 --- Ping Pong with Channels

Projeto de concorrência usando goroutines e channels.

Saída esperada:

ping pong ping pong

Conceitos: - goroutines - channels - sincronização - bloqueio de
execução

------------------------------------------------------------------------

# Projeto 2 --- API da Vovó \[api_gorila\]

API REST simples com CRUD de usuários usando Gorilla Mux.

Rotas: - GET /users - GET /users/{id} - POST /users - PUT /users/{id} -
DELETE /users/{id}

Execução:

cd api_gorila go run .

Servidor: http://localhost:8080

------------------------------------------------------------------------

# Projeto 3 --- Testes Unitários \[calc_test\]

Projeto para prática de testes automatizados com package testing.

Arquivos: - calc.go - calc_test.go

Funções: - soma() - multiplica() - divide() - subtracao()

Execução:

cd calc_test go test

Detalhado:

go test -v

Padrão usado: AAA - Arrange - Act - Assert

------------------------------------------------------------------------

# Estrutura

README.md api_gorila/ calc_test/

------------------------------------------------------------------------

## Autor

Projetos criados para estudo de Go Backend, APIs REST, concorrência e
testes automatizados.

Happy coding 🚀
