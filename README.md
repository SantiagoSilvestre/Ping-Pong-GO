# Go Concurrency Example: Ping Pong with Channels

## 📌 Overview

This project demonstrates basic concurrency in Go using:

- Goroutines
- Channels
- Synchronization through blocking operations

The program prints:

```
ping
pong
ping
pong
...
```

Continuously, in a controlled alternating sequence.

---

## 🧠 How It Works

### 1.Ping & Pong

Two goroutines continuously send messages to their respective channels:

```go
func ping(c chan string) {
	for {
		c <- "ping"
	}
}

func pong(c chan string) {
	for {
		c <- "pong"
	}
}
```

- `ping` sends `"ping"` to `c1`
- `pong` sends `"pong"` to `c2`
- Both run infinitely

---

### 2. Imprimir

```go
func imprimir(c1, c2 chan string) {
	for {
		ping := <-c1
		fmt.Println(ping)
		time.Sleep(time.Second * 1)

		pong := <-c2
		fmt.Println(pong)
		time.Sleep(time.Second * 1)
	}
}
```

This function:

1. Receives from `c1` → prints `"ping"`
2. Waits 1 second
3. Receives from `c2` → prints `"pong"`
4. Waits 1 second

👉 This enforces **strict alternation**

---

### 3. Main Function

```go
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go ping(c1)
	go pong(c2)
	go imprimir(c1, c2)

	var input string
	fmt.Scanln(&input)
}
```

- Starts all goroutines
- Keeps program running until user input

---

## ⚠️ Important Notes

### Blocking Channels

Channels are **unbuffered**, meaning:

- `c <- value` blocks until someone reads
- `<-c` blocks until someone writes

This ensures synchronization between goroutines.

---

### CPU Usage Consideration

The producers run infinite loops without delay:

```go
for {
	c <- "ping"
}
```

This is safe because the send operation blocks, but in real-world systems, it's better to use:

- `time.Sleep`
- or `time.Ticker`

---

## 🚀 Improvements (Next Steps)

You can evolve this example by:

- Using `time.Ticker` for controlled intervals
- Adding `context.Context` for graceful shutdown
- Using `select` for multiplexing channels

---

## 🎯 Key Takeaways

- Goroutines run concurrently
- Channels synchronize communication
- Order must be controlled explicitly if needed
- Go concurrency is simple but powerful

---

## 📦 Output Example

```
ping
pong
ping
pong
...
```

---

# API da Vovó [api_gorila](/api_gorila/)

## Descrição

Este projeto é uma API REST simples desenvolvida em Go utilizando o
pacote Gorilla Mux.

Ela permite realizar operações básicas de CRUD de usuários:

-   Criar usuário (POST)
-   Listar todos os usuários (GET)
-   Buscar usuário por ID (GET)
-   Atualizar usuário (PUT)
-   Deletar usuário (DELETE)

Os dados são armazenados em memória utilizando um slice
(`var Users []User`), ou seja, não existe banco de dados neste projeto.

------------------------------------------------------------------------

## Tecnologias utilizadas

-   Go
-   Gorilla Mux

Pacote utilizado: github.com/gorilla/mux

------------------------------------------------------------------------

## Como instalar

### 1. Instalar o Go

Verifique se o Go está instalado com:

``` bash
go version
```

Se não estiver instalado, faça o download no site oficial:
https://go.dev/

------------------------------------------------------------------------

### 2. Criar o módulo Go

Dentro da pasta do projeto execute:

``` bash
go mod init api_gorila
```

Isso cria o arquivo `go.mod`.

------------------------------------------------------------------------

### 3. Instalar a dependência

Execute:

``` bash
go get github.com/gorilla/mux
```

Isso instalará o pacote necessário para as rotas da API.

------------------------------------------------------------------------

### 4. Organizar dependências

Execute:

``` bash
go mod tidy
```

Isso atualiza e organiza o `go.mod` e o `go.sum`.

------------------------------------------------------------------------

## Como executar o projeto

Na pasta onde está o arquivo principal, execute:

``` bash
go run .
```

ou

``` bash
go run nome_do_arquivo.go
```

Exemplo:

``` bash
go run api_gorila.go
```

A API será iniciada em:

http://localhost:8080

------------------------------------------------------------------------

## Rotas disponíveis

### GET /

Página inicial da API.

Exibe mensagem de boas-vindas e lista das rotas disponíveis.

------------------------------------------------------------------------

### GET /users

Lista todos os usuários cadastrados.

Exemplo:

http://localhost:8080/users

------------------------------------------------------------------------

### GET /users/{id}

Busca um usuário pelo ID.

Exemplo:

http://localhost:8080/users/1

------------------------------------------------------------------------

### POST /users

Cria um novo usuário.

Body JSON:

``` json
{
  "id": 1,
  "nome": "Santiago",
  "email": "teste@email.com"
}
```

------------------------------------------------------------------------

### PUT /users/{id}

Atualiza um usuário existente.

Exemplo:

PUT /users/1

Body JSON:

``` json
{
  "nome": "Novo Nome",
  "email": "novo@email.com"
}
```

------------------------------------------------------------------------

### DELETE /users/{id}

Remove um usuário pelo ID.

Exemplo:

DELETE /users/1

------------------------------------------------------------------------

## Como funciona internamente

A aplicação utiliza:

``` go
var Users []User
```

para armazenar os usuários temporariamente.

Isso significa:

-   não existe banco de dados
-   ao reiniciar a aplicação os dados são perdidos

Este projeto é ideal para estudos de CRUD, rotas REST e aprendizado de
APIs em Go.

------------------------------------------------------------------------

## Estrutura principal

### Struct User

Representa um usuário:

-   Id
-   Nome
-   Email

### Funções principais

-   createUser()
-   updateUser()
-   getUserById()
-   deleteUserById()

Cada uma representa uma operação do CRUD.

------------------------------------------------------------------------

## Observação

Este projeto é uma base inicial.

Melhorias futuras:

-   integração com banco de dados
-   uso de PostgreSQL ou MySQL
-   separação em camadas (handlers, services, repositories)
-   autenticação JWT
-   validações avançadas
-   retorno em JSON padronizado

------------------------------------------------------------------------

## Autor

Projeto criado para fins de estudo de APIs REST com Go.


Happy coding 🚀
